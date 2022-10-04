// Created by EldersJavas(EldersJavas&gmail.com)

package qhash

import "testing"

func TestBLAKE2b_256(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "8b7ca7d27d9fc55fa30abfe515b3afb24e3fe89fdd02e2ac92bca2c96680642e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := BLAKE2b_256(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("BLAKE2b_256() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestBLAKE2b_384(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "bc3868e9afae58a3c910c16ff94d1a6185e8c588f77132979395f9d01126aafe32770cfb33a8a84dddd0bf0c826cb9eb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := BLAKE2b_384(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("BLAKE2b_384() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestBLAKE2b_512(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "ef15eaf92d5e335345a3e1d977bc7d8797c3d275717cc1b10af79c93cda01aeb2a0c59bc02e2bdf9380fd1b54eb9e1669026930ccc24bd49748e65f9a6b2ee68",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := BLAKE2b_512(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("BLAKE2b_512() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestBLAKE2s_256(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "f73a5fbf881f89b814871f46e26ad3fa37cb2921c5e8561618639015b3ccbb71",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := BLAKE2s_256(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("BLAKE2s_256() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestMD4(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "a58fc871f5f68e4146474ac1e2f07419",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := MD4(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("MD4() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestMD5(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "8b1a9953c4611296a827abf8c47804d7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := MD5(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("MD5() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestMD5SHA1(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "a58fc871f5f68e4146474ac1e2f07419+f7ff9e8b7bb2e09b70935a5d785e0cc5d9d0abf0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := MD5SHA1(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("MD5SHA1() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestRIPEMD160(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "d44426aca8ae0a69cdbc4021c64fa5ad68ca32fe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := RIPEMD160(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("RIPEMD160() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "f7ff9e8b7bb2e09b70935a5d785e0cc5d9d0abf0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA1(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA1() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA224(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "4149da18aa8bfc2b1e382c6c26556d01a92c261b6436dad5e3be3fcc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA224(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA224() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA256(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA256() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA384(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "3519fe5ad2c596efe3e276a6f351b8fc0b03db861782490d45f7598ebd0ab5fd5520ed102f38c4a5ec834e98668035fc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA384(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA384() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA3_224(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "4cf679344af02c2b89e4a902f939f4608bcac0fbf81511da13d7d9b9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA3_224(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA3_224() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA3_256(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "8ca66ee6b2fe4bb928a8e3cd2f508de4119c0895f22e011117e22cf9b13de7ef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA3_256(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA3_256() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA3_384(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "df7e26e3d067579481501057c43aea61035c8ffdf12d9ae427ef4038ad7c13266a11c0a3896adef37ad1bc85a2b5bdac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA3_384(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA3_384() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA3_512(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "0b8a44ac991e2b263e8623cfbeefc1cffe8c1c0de57b3e2bf1673b4f35e660e89abd18afb7ac93cf215eba36dd1af67698d6c9ca3fdaaf734ffc4bd5a8e34627",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA3_512(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA3_512() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "3615f80c9d293ed7402687f94b22d58e529b8cc7916f8fac7fddf7fbd5af4cf777d3d795a7a00a16bf7e7f3fb9561ee9baae480da9fe7a18769e71886b03f315",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA512(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA512() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA512_224(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "0d075258abfd1f8b81fc0a5207a1aa5cc82eb287720b1f849b862235",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA512_224(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA512_224() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}

func TestSHA512_256(t *testing.T) {
	type args struct {
		UseByte []byte
	}
	tests := []struct {
		name         string
		args         args
		wantHsResult string
	}{
		{
			name:         "",
			args:         args{[]byte("Hello")},
			wantHsResult: "7e75b18b88d2cb8be95b05ec611e54e2460408a2dcf858f945686446c9d07aac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHsResult := SHA512_256(tt.args.UseByte); gotHsResult != tt.wantHsResult {
				t.Errorf("SHA512_256() = %v, want %v", gotHsResult, tt.wantHsResult)
			}
		})
	}
}
