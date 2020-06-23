// Code generated by "esc -private -pkg=layers -o=bindata.go trans.hsaco gpu_gemm.hsaco relu.hsaco maxpooling.hsaco"; DO NOT EDIT.

package layers

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/gpu_gemm.hsaco": {
		name:    "gpu_gemm.hsaco",
		local:   "gpu_gemm.hsaco",
		size:    13800,
		modtime: 1591217294,
		compressed: `
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
/xcAAP//aVmkveg1AAA=
`,
	},

	"/maxpooling.hsaco": {
		name:    "maxpooling.hsaco",
		local:   "maxpooling.hsaco",
		size:    18384,
		modtime: 1592882172,
		compressed: `
H4sIAAAAAAAC/+xcXWwbV3Y+c4YcDkfDmSE54p8omWJoy6EtmRoxCtctuvFPrQSxE2e9zW66DVxapGTa
FClQlGMXq1lStSTaCFDVWGQXCyMsgqDNw/ahQB/6Uv0UxT7sLhamgEVf/FC0XbToY/+QvpDFnbkjkYRp
J0gW2cq8gHjE755z7jnn3jnn4s5wvve7Fy8gw7wCtLHwj8CQfxLmd6vj0XmTxg0sBTy8AiIIwAGArY2v
m+4wnZSnOEPlerWP450UlAM5e5t93ZSXO2m7HAe9DV2ETmrJ4eeUs/z7xq/KGdtnkGu3j7S3flXOcPD5
m82KJ8KB4W3007FOamuT4+n4Zy6dN9ituRky1oOJ28Cx75uFnbl0fuby75m8QQAYoHh6ITM/WxhPL2TI
3/Wl9Pj4/NztVGKK6r01BiBQ3vHxceHtbGkpVyycjljtO5HJk5FE5F3h9WypkM0vnRYikfHIG+mF7AFP
JBK5lL59uVjMXyiW3kuXMgKBrtxZuFbMt3GOdTK9cjMzZjBeTBfml9PzBwrfXMwWzl2MnOvo3bfMsEiL
vGv0ninNGyaR9gSzCuXrpWw6syRYwDfvLGY7uHKF8n7nldwfdYon97vO5HPzhdNP7Ho7nV/Ovp4rZKzu
s3cMqJOBDGwxvDalHSienX1rOZ0/UH0+O5dezpd7+3StWC4XF65m0uV0b7fG5vLFdDk+1tu3VG/fUr19
m8kXr6XzZ5fn5rKl3g5eaHcwkyldWUzPZi03TRVfIACF5YXDNJ+z19MFcmUdJp+uZ3Pz18uHyaP3cpny
9cPk0GKxmM9mrh6+maKOHboJu2kU4KuH0af3DpNPS+VSLpM9XPNEfTpU87SYzhyuSSIOHaoZKhcXn+Nt
7kJ66eaz3M8VvjLnX/u1Op8rZLK3l54Dz78ky1/NZTLZgjn4m3NzS9nyt5/iwHTy1z/+O1/x+L//FYz/
RrHwtGSY+ozLpm/VV2bVpeV8OTdTymWu3CnMninNf2ELzxUz2cul4uL+KdXr2VIhXZq/kp1fyBbKpvGT
SWtBzpSKy4u070LudjZjMiRo9+VS7la6nO3N0Kmd+m9Z+q30rexcqWiNGonsXwdvLC9cmbn8jaX9WCXb
et7u6NGsoS6lb1/Ip8vfKpZumlYbSrWXpp96Qng2PXvz2UeEFlf/jPAL7Jxyc3PP586JeE92T8/jxql/
ONo/HO0fjvYPR/uHo/3D0f7haP9wtH84ekgPR63HAJ6XXX7/oKx/UNY/KHveD8qmpn/DDsqmUj0Pyl5+
5kHZxMSE0Ps5QgYAQiwH8DJ9npE+9Oe2cPqc5d9wB/zWX8h8nJGL0O/n/2I2850jux+w9JlBix+7B+06
gIPOZ/ag3/qt3/qt3778ZuVjxni6mz14EL1HO8/8GDaNZ707i8hm2/8nQe18Nt1m41qtVus30X+WlXYU
Iwa4YwOACqLOo03/Hry/DS1cN63+twba0OYS4QEo/APezT/gobZyI1LZ+gOmto0i6i24twVQefPvAHdE
otcu7JBSyTq5HSOynLDzCqGMsHOZxHsGdVkafVBhPtpyIgqMjDrC2h5jJ7OwbtCKjU8C/PRRhUMyRw10
Erq+hywCspyGLE4zAA3GgUDsZYicg/hi6qkip+0ir7E2IcnaHq4A/OwRCghwo7JuydmILspve5+r2CuO
VeW+UGVtP1ypIKbc9weqrvti1dYKVy7DgW57ZXjV1hqp/CEAxJDT7ZUjq0eR12VY3dZjzQ0dmxt2+JO9
VUCoSpI2hKiHBEFjXGIK7KDxsL5yI/rRBsDiY7sDQI827zmZ2ra7DptuYKoMYI3EtQ5iFERPXKnHNknM
KijodgzoVswc1B7eipmkmjFTEBwADccgoet7DhYBXIrGu8S4g1U0hxU7N4IRNxeCEoMqxI7WjoqSzlO9
u6hqjhlJ59y+JOemMfQhcBhsVLEtli4EpyQKnCzqAovAeDHOeEeSDKztca+DYQOh9RElWkGPaaNExoYG
4z2wQXZLmsPticcUJc64JI1xUTtlBIHweBFc9dgmsKLOwZ/ufSggDMDankB0A0ad/lDyaOiDFRcEG7sq
gmvfp7AO8ItHTj+Ck9iuItyAf1+XZny64PWAE6Dh8iI43Z6406cKgqzqTncoWQ/5ojzV74R1g1aUsGm/
B0nWaPDDaMSfH0Rw+T2a0++J84NhjR9Upq1+YjvvRbCJoibKUorMO9lCcijqDjTjTXQbOAJwsLYnvwCw
qwY0p1tKOt0PV3gINcBFstXPHh2V6Byo5hyQdeAaNX6r0xBcCKJLhJjPE+e8Ypzz+pKy7EkSnc4XyR51
bW9gCEDAgC5iWPeJsekwirprNJocwJA+jJIO8N3HnEjmbW0PBhiws6o+goJ+BEUdoPzY7gRYkwStLohR
Y50LCHaUdJsgaA6XmHKzqqHDzgO4RwXNCWt7owKCJyZokiAEKjExVPeoUSeR9SDsoqSzXk6zDfIpbmB1
5UZkdWuiVdsm4/MnyPpe25NFhArz8dZ/Ykj3S9K0V/JoFUkNyR5PQPCurgxAsCG4TV0iuQ6PNTc4CDaq
IpmfYENWSOx+8ogbQIDYxxuEX/Yh6LHmPU+rtn0j8udb/9Gqbdd9gahE7PKR8T7ZkiDYYIcRZAg3bEcQ
BDagA1x+LAF4AhjRgxjVvRBs8EEEkYwXMMcRvQhw/JONRZJLYfGx4iHza8gpROfdsKmzMoJg8px/pAwh
uMhaHKXX7xiCfrx5759b97b1482Nn7ZWtwGq2wCVL+Xv89Y/rkf9q7X9f8z45dT/j/pn1T2z3oFR73iA
ByziDuk/B2jUK7O+kLxMcrdZ+/6KqW2zIJj1jZV2IrTukTr3xwy38yrRz5j10JBHTnfVnZuANt3lhKpY
D2yyKOjg9NXEAFRlyf6gwtzfQuT1FvzAqKcQCNaMmoLNDT3U3CA5jqN1xYGoc6KoMbKUsgGtKcJ9o6Zw
AoAuNO/VmNq2z6ojLlULipJO6gjJ6X5aQ/ySJDhkkoPW95hxkufWDGrVkYrHlyR5jOROI8+xCILLE3eK
YpxnVY23askgzeNeNK5XoqOKqraLPs3pDhg5jOQuX6CtbgwiDJB64cI441KMeuGYInls3aB1xROtIK1p
Iq0X1jgyAu8VNZLjAh5PnJFVjZG7bHGhkcMYolNWk0H1gxWB1AZSu+AXjzj5oJaRekD4ifxBDZNMm4IA
PCsmVTGgKYF1Y/51+FFtHGrbnjpuehCqgGyNFUmNjhkxNuMb0pUATtfRE2W8qMkkTztIDl43KBkDPZ54
ZSRq+hghtQIa8gtmzRCGEIQhnyYMqdMEF/wUZ2ndcwBU1bC2qx7TGP/xJOM3Y8weP4ixMIzg3fcHk1VE
LTRa2WJR1RF9ug7NmrLvB1tl2/wg/YSv3R8Pi3q3H8R+HZobZN75IYB6IBwluVAcRgC3J1mJfH9FfkEx
ag9/lK6jIQRxeCROcjM/auZOl5v0hUwaVAz9/BGEQU9Ac3fF3VvHTe9+3BWdn1F0v6IIvBzS3QFluq7E
o/wJRXNTWz3UVv6EEsd4PF6JHjNjHkOyY2+4xwhd31OHENShk5o6dHya4Oo4go/gboQKHtf9sLbnI3E/
mdB2T2oafyqZ5E/RuCfNuJPrxTeJEHQrxnj8CSVZVRQa97iOeHI/7qYfVtwVnYxB+glfu08qq+jdvhAf
SNyJTcogguT3xeVT4aQ0+NBc5xFzncunECSyziPmOleGw8Y6I2tD8Ktxwe9LCrC2Z5cQqr6wJvlqKxJK
el31RQWSb1QEx7AEOGjTWD+Xsg3ymt0vpBh5daWCqn4j8qOtf2jWtiux76+4x3wgkT3Gi2SNru8NDCHU
A6q5HgYRxMFAnNTegWEEUgddfsIXMumwz5QJIgjjapzsWbghABfRdwJBHAonpRMPV4jcrs/0TSTriPjm
M30bCBId0JCIvkFffGBQTVZVnyartRXio4wIMtn7+H3A+1dXiF1k70b2CpyX1PzVLeLTAMm7I80NPdzc
YIDud+Enj0jeAVzdIHJyAEHH5r0fNMn+4eHWf5H9QyAUlUnMAkTXh1tkL4pHiK/hBjtKan7Q2EcoEG7Y
RxF41qMrZKxoc8NJxvGQPUCwcTdkylTC5rgk/8LIhxtE/u6IKU/yhT7SvPfLVm2b7C1kADfZS7hJfKJo
5HA32TNGPtr6V8rjojzCCyr+0/5+4vls1m/NP33RpAP0e4BSO6V36e/wrV1PhNL/braKxs/a452/K8/E
nzzebD5dmI/cMp+ZjExOTiQmEpHjp5ZKs6eyt8vZUiGdP5XP31oYXywVb2Rny6dMgcRUaiqdmZydu5ZI
p7KJl64l0gktkU1NT2nJRGLya8lpLXVNS069CHAxV7iZLZ2OXLx4/rPoz+czn0f7085RSDT/eroT5yn+
Ta0Tlyj+7S5cpfhmFx6i+OOXOvEIxcWucY9RPNWFn6T44m934lMU//su/DTFI7/TiZ+heKwLf5XiF7vw
Nyn+fhf+NsV/3oW/S/FPUl3rh+K2r3fiNyj+4y7+RYr/Sxd+i+Kfnu7Ev2vZ/1ud+Kp1tbzcif/QwhOd
+M+NzbXj4L0PtP1lj/srf9vj/gpMFIrlLExk7hSW7izAxHxheeJ6euk60E+Cl0swUc7eLhvf0gu5WZiY
LS4sZAtlmFi6s1BOX4OJpetL5ZL5n0nh7NnE1clJg2jG55TxmTQ+p43Pl43Pr8HZs5NXJxMm0UySNMlL
Jpk2ScokpoDJqE0aZMr4NIVM5lT3XZ9Yvjibznfd++kEe90gunr+nTfOXHrt3Jd0Hupov0fV4z0a7e/d
aG8OmkOxK69a9NW2vMq0vS/EyrcyAPxPq1W05K28atFYl1l81/hBqhu78rBFI13yti56hN6rw668f7dL
vjPvHbSx9netPOX9LL0UjFNZ635hr/em2Lv8p69RgWmqsutyhEUqv91jeIt+/Un3J4m+M1QPHtTJwBPm
b6bd9rb2Z3RxvvOM+L3VQ/5/qfwvnyH/fwEAAP//ycy1bNBHAAA=
`,
	},

	"/relu.hsaco": {
		name:    "relu.hsaco",
		local:   "relu.hsaco",
		size:    13904,
		modtime: 1591217294,
		compressed: `
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
/afZ/xoAAP//bciEz1A2AAA=
`,
	},

	"/trans.hsaco": {
		name:    "trans.hsaco",
		local:   "trans.hsaco",
		size:    9656,
		modtime: 1591217294,
		compressed: `
H4sIAAAAAAAC/+xaTXPbxBt/JCuO4+T/B0605YAIh7QdJCtynDi50Lw0SadJmjbQUphOZyOtbTV6G2kd
YgbSF2inM3SGl+FePgCHHjm1OfABmJ576IHO8AXgiBlJu7bkWn2BdgpUv5n4sZ6XfX7PWqvsrvbi0ZVF
nuOOAEUO7gMXfHk1umaGH0cjeTjUVaEAR2AEipAHACHm1yv3uKQsUD1H49Lw9khSMj5B3EDsulf+yiVl
PC7gCiLV90gXkpLF8U8Zx+o79YDowhPExfkFOPmA6Hl4egisP3noEo/JW0NJKcTiCjT/7OpC6M5+mzfC
+yHSCzDYqY3pZlcXltbfj3z3A8Aw1SNLr2u2hCw9+Gv4SJLqtZ2qUqbtykMAReorSVLxNPZ8w7FnRIaP
xPF3REU8VzyOPRub/kxRFCVxDVm46yOK4nsesn3X8XExuNpoWZuOGXMa69iPbOljoc8KsutNVO82c8LF
9vyKOJ+wdviEPFTxXGid9eohkQB9yByzix1eLRcn7GM100Hk8FjHY8P4JBld7ZhmTaNuz/Q1nUZmEx83
bJ2Zl0xnE5lzzVoNe0mvgALzWiyr3dZ13dtwkYZPNpE502mia9c0ZomwgGuoaZL0uk80yctZuNMkbpOc
MXTSSO8AwybptU+k1z6RXvtcK1Sll30sXvZfLGsZG/UG+RfX9YxutGVD17Ed3SsnajUfkw8eQXBy4vnn
P/uC83/4AvKvOfaj7ovqE47yjNULY7XaNImx5Bn6RsvWZr3632Y47+h43XPczr/kYKqAvPoGrlvYJhH5
qkKNS57TdKlp0djBemRn5nXP2EYEpzskG6flM6Jn0DaueQ5LKoqdYbDWtDaW1k/5na4qV7qW0wkLM6yi
nUUTkTOOtxWRDttUK5NFWZaL6fO/YM52IJd/aB7Mxf4O0OmmSK+PTg9/x8emily8wc7kCTJkyPAfBUfH
Pxeu7nI9D4E+/twP8PVQsNZLPozW42v3aIXeXZsKQr7dbrf/ifXzwO8Fa9LPIb8XPBtzueLePgjWp1/e
Hobvb1+EG3egDdcC9gUofDtSKFwWrsGlgUvcFY7P7wIv7AIs3+MB+AKACvDTzzzw4MKNOzxcvTsk8PB/
QVD5HK/m+OufwYVL10D84vYBuB7aYRAgxxX3wr4dENSbQn6Uh6/uXhF4CHIP8IVdQRBUYTBfBVgP83Bw
9S43ADDI53c5DtSbwI9CEAM85AHUXI6vBvkB3HvRw/3ynexOz5AhQ4YMGTJkyJAhQ4aXG+xd863hSFIB
+6hkK/l99D08W/Uzv9/+aDuBvE8VnXf6I/3zLc/Pi5qJ7Lq4Hb1lFccVWZEV8WBJRwSVLmB7y7B96WPH
2/JdpOGS5lhuk2DJczRL8rAplWWlhHcI9mxklhqaJhGHlExz25Jcz7mANVKKEkwoU5XJ8uTE1Pg0rpaR
qut6BW1iBdVqeEpRptSypuvVinpIPLiJfKyLji0G9MqyIo9PT0yXpakKRtMVVaItiYdgxbC3sDcjrqws
PA/ipqk/Oe3H7usEv+6nryX1g1T/S4/+9XAzZLB7ToHifyn7yiDbDsEg6y3bb1kg1+2m3EB+A+hnoCce
yATvkPAKWYYGsuZYFrYJyH7LImgTZL/hEy/6FkmYm1POj4efamwT+vzC2bXZ1WPzz3LfazC29512XqKz
hwUP9+NwLIyNHyaV2PjhYudC2Lh6BQB+b7cdFs/GD5NiD61CT/79tG2+d7yNJPPwPfyZfJOe4+B7xjeT
w33vpy7G4mdqIP0cTloDEo3NMUXK+ZiBnvpZmknapNKTxqXxq1z/9Ey+G++7GPbeiuRO7Dkn9Pn9luLc
Y6jSc1JnH9N/J1PivxlNdkda/J8BAAD//5VIJrK4JQAA
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
