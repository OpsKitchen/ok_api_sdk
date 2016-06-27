# ok_api_sdk_go
Golang SDK for OK OpenAPI platform

## Get started
### Prepare Go SDK
Download go SDK from: https://golang.org/dl/

Decompress it into /usr/local/go

Run:

    export PATH=$PATH:/usr/local/go/bin
    export GOROOT="/usr/local/go/"


### Download or clone this project:

    mkdir ok_api_sdk_go
    cd ok_api_sdk_go
    git clone https://github.com/OpsKitchen/ok_api_sdk_go.git src

### Customize your demo data
Edit src/example/simplest.go, replace the demo data

### Run
    cd src
    ./build.sh
    ./example

Demo output is like this:

    [root@dev67 example]# go run simplest.go
    [DEBUG]: Gateway url: http://api.OpsKitchen.com/gw/json
    [DEBUG]: Request body: api=ops.meta.os.list&version=1.0&timestamp=1466961052&params=null
    [DEBUG]: Request header: map[Oa-App-Market-Id:[678] Oa-App-Version:[1.0.1] Oa-Device-Id:[08:00:27:bf:d4:5e] Oa-Session-Id:[] Oa-Sign:[81129ce782af9d0ce1f8b5419d00ab0b] Content-Type:[application/x-www-form-urlencoded] Oa-App-Key:[101]]
    [DEBUG]: Response body: {"success":true,"data":[{"id":"3","name":"centos","pms":"yum","repoUrlSample":"http:\/\/ok.com\/centos\/$releasever\/$basearch","fullName":"Community ENTerprise OS","description":"i386\/x86_64","homepage":"","rank":"3","enabled":"1"}]}
    &{0xc82000baa0   [map[enabled:1 id:3 name:centos pms:yum fullName:Community ENTerprise OS repoUrlSample:http://ok.com/centos/$releasever/$basearch description:i386/x86_64 homepage: rank:3]]}
    [DEBUG]: Gateway url: http://api.OpsKitchen.com/gw/json
    [DEBUG]: Request body: api=ops.meta.osImage.listByOsReleaseId&version=1.0&timestamp=1466961052&params={"osReleaseId":"3022"}
    [DEBUG]: Request header: map[Oa-Sign:[3b89f21740a52dc4b750af96aee05c7d] Content-Type:[application/x-www-form-urlencoded] Oa-App-Key:[101] Oa-App-Market-Id:[678] Oa-App-Version:[1.0.1] Oa-Device-Id:[08:00:27:bf:d4:5e] Oa-Session-Id:[]]
    [DEBUG]: Response body: {"success":true,"data":[{"id":"302207","osReleaseId":"3022","fullVersion":"6.7","repoSnapshotId":"307"},{"id":"302206","osReleaseId":"3022","fullVersion":"6.6","repoSnapshotId":"306"},{"id":"302205","osReleaseId":"3022","fullVersion":"6.5","repoSnapshotId":"305"},{"id":"302204","osReleaseId":"3022","fullVersion":"6.4","repoSnapshotId":"304"},{"id":"302203","osReleaseId":"3022","fullVersion":"6.3","repoSnapshotId":"303"},{"id":"302202","osReleaseId":"3022","fullVersion":"6.2","repoSnapshotId":"302"},{"id":"302201","osReleaseId":"3022","fullVersion":"6.1","repoSnapshotId":"301"},{"id":"302200","osReleaseId":"3022","fullVersion":"6.0","repoSnapshotId":"300"}]}
    &{0xc8200f01b0   [map[id:302207 osReleaseId:3022 fullVersion:6.7 repoSnapshotId:307] map[id:302206 osReleaseId:3022 fullVersion:6.6 repoSnapshotId:306] map[osReleaseId:3022 fullVersion:6.5 repoSnapshotId:305 id:302205] map[fullVersion:6.4 repoSnapshotId:304 id:302204 osReleaseId:3022] map[id:302203 osReleaseId:3022 fullVersion:6.3 repoSnapshotId:303] map[id:302202 osReleaseId:3022 fullVersion:6.2 repoSnapshotId:302] map[id:302201 osReleaseId:3022 fullVersion:6.1 repoSnapshotId:301] map[id:302200 osReleaseId:3022 fullVersion:6.0 repoSnapshotId:300]]}
