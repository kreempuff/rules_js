import { create, Archiver } from 'archiver'
import * as path from 'node:path'
import { readdir, readFile, realpath, stat } from 'node:fs/promises'
import { createReadStream, createWriteStream } from 'node:fs'

const [entriesPath, appLayerPath, nodeModulesLayerPath] = process.argv.slice(2)

const app = create('tar', { gzip: true })
app.pipe(createWriteStream(appLayerPath))
const app_structure = new Set<string>()

const node_modules = create('tar', { gzip: true })
node_modules.pipe(createWriteStream(nodeModulesLayerPath))
const node_modules_structure = new Set<string>()

const entries: { [k: string]: string } = JSON.parse(
    (await readFile(entriesPath)).toString()
)

function mkdirP(p: string, output: Archiver, structure: Set<string>) {
    const dirname = path.dirname(p).split('/')
    let prev = '/'
    for (const part of dirname) {
        if (!part) {
            continue
        }
        prev = path.join(prev, part)
        if (structure.has(prev)) {
            continue
        }
        structure.add(prev)
        output = output.append(
            null as any,
            { name: prev, type: 'directory' } as any
        )
    }
}

function findKeyByValue(value: string): string {
    for (const [key, val] of Object.entries(entries)) {
        if (val == value) {
            return key
        }
    }
    console.error(`couldn't map ${value} to a path.`)
    process.exit(1)
}

async function* walk(dir: string, accumulate = '') {
    const dirents = await readdir(dir, { withFileTypes: true })
    for (const dirent of dirents) {
        if (dirent.isDirectory()) {
            yield* walk(
                path.join(dir, dirent.name),
                path.join(accumulate, dirent.name)
            )
        } else {
            yield path.join(accumulate, dirent.name)
        }
    }
}

for (const key of Object.keys(entries).sort()) {
    const dest = entries[key]
    const outputArchive = dest.indexOf('node_modules') ? node_modules : app
    const structure = dest.indexOf('node_modules')
        ? node_modules_structure
        : app_structure

    mkdirP(key, outputArchive, structure)

    const entryStat = await stat(dest)
    const realp = path.relative(process.cwd(), await realpath(dest))

    if (entryStat.isFile()) {
        outputArchive.append(createReadStream(dest), {
            name: key,
            stats: entryStat,
        })
    } else if (entryStat.isSymbolicLink() || dest != realp) {
        outputArchive.symlink(key.replace(/^\//, ''), findKeyByValue(realp))
    } else if (entryStat.isDirectory()) {
        for await (const sub_key of walk(dest)) {
            const new_key = path.join(key, sub_key)
            const new_dest = path.join(dest, sub_key)
            mkdirP(new_key, outputArchive, structure)
            outputArchive.append(createReadStream(new_dest), { name: new_key })
        }
    } else {
        console.error(`unknown type ${entryStat}`)
    }
}

await app.finalize()
await node_modules.finalize()
