# markymark

![markymark](./markymark.png)

## Overview

Markymark, or `mm`, is a tool that you can use to publish, share, download
[Markdown](https://en.wikipedia.org/wiki/Markdown) (.md) documents over [OCI Distribution](https://github.com/opencontainers/distribution-spec).

## Getting started

Requires a valid Go 1.12+ environment.

Clone this repo and run the following:
```
make bootstrap build
cp bin/mm /usr/local/bin
```

This will give you the `mm` CLI:
```
mm --help
```

Note: the examples below assume an OCI Distribution-compliant registry running at `localhost:5000`.
To quickly start a local registry using a [docker/distribution]() server, run the following:
```
make start-registry
```


## How to use

The `mm` CLI attempts to be Unix-y by being interoperable
with other tools via pipes and redirects.

### Pushing Markdown

Markdown document content is read from stdin:

```
$ cat README.md | mm push localhost:5000/md/readme:master
Pushed localhost:5000/md/readme:master
Size: 2.6 KiB
Digest: sha256:c0fc0876711655a6e90d621015d67f0072d764cb24df88b1ba327d70674ff54a
```

```
$ echo "# Hello world" | mm push localhost:5000/md/hello:0.1.0
Pushed localhost:5000/md/readme:master
Size: 14 B
Digest: sha256:75a033b1054ea103961566fd597658834832efaa194de7a4a18771fd8506cfa6
```

### Pulling Markdown

On pull, Markdown document content is printed to stdout (and logs to stderr):
```
$ mm pull localhost:5000/md/hello:0.1.0 > hello.md
Pulled localhost:5000/md/hello:0.1.0
Size: 14 B
Digest: sha256:941ce5c4eee6f77a2552f7cf7acf43916cc965422f1ccde7a208e680898f711b
$ cat README.md
# Hello world
```

## Manifest

To examine the manifest created by `mm`, you can inspect a registry's storage backend.

Example:
```
$ cat data/docker/registry/v2/blobs/sha256/8a/8ac3cf7b8df11ad743a1c1aea82a98af812014a44353f556454b3205105977ba/data | jq
{
  "schemaVersion": 2,
  "config": {
    "mediaType": "application/vnd.daringfireball.markdown.config.v1+json",
    "digest": "sha256:44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a",
    "size": 2
  },
  "layers": [
    {
      "mediaType": "application/vnd.daringfireball.markdown.content.layer.v1+md",
      "digest": "sha256:2bb75fb61901222362eba5a42fddb31a69562442cde19999b0e2c24fdaf09644",
      "size": 2689
    }
  ]
}
```

## Mediatypes

Markymark introduces two (2) new OCI mediatypes:

1. Manifest config mediatype: `application/vnd.daringfireball.markdown.config.v1+json`
2. Content layer mediatype: `application/vnd.daringfireball.markdown.content.layer.v1+md`
    
If you wish to use this tool with an existing registry,
keep in mind that the registry must be able to accept these mediatypes (e.g. via whitelist).