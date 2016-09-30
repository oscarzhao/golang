package mail

import (
	"testing"
)

func TestDecode(t *testing.T) {
	shouldSuccess := map[string]string{
		"=?utf-8?q?=E5=93=88=E5=93=88=E5=93=88=EF=BC=8C=E4=BD=A0=E5=80=92=E6=98=AF?= =?utf-8?q?=E6=9D=A5=E6=89=93=E6=88=91=E5=91=80?=": "哈哈哈，你倒是来打我呀",
		"kill me now":                             "kill me now",
		"=?UTF-8?Q?=E6=9D=A5=E8=87=AA=EF=BC=9A?=": "来自：",
		`<p>=E6=9D=A5=E8=87=AA=EF=BC=9A694920142@qq.com</p><p>=E7=94=A8=E6=88=B7=
=E6=9C=80=E6=96=B0=E5=9B=9E=E5=A4=8D=EF=BC=9A<p>=E6=88=91=E4=B8=80=
=E5=85=B1=E6=9C=8950=E4=B8=AA=EF=BC=81=EF=BC=8C=E4=B9=8B=E5=89=8D=
=E5=85=91=E6=8D=A2=E4=BA=8616=E4=B8=AA=EF=BC=8C=E8=BF=98=E6=9C=8934=
=E4=B8=AA=EF=BC=81=E5=BF=AB=E7=BB=99=E6=88=91=E5=85=91=E6=8D=A2=E5=91=80~</=
p>
</p><p>=E5=B7=A5=E5=8D=95=E5=86=85=E5=AE=B9=EF=BC=9A<p>=E4=B9=8B=E5=89=8D=
=E5=85=91=E6=8D=A2=E8=BF=8716=E4=B8=AA=EF=BC=8C=E7=8E=B0=E5=9C=A8=
=E8=BF=98=E6=9C=8934=E4=B8=AA=EF=BC=8C=E7=BB=99=E6=88=91=E5=85=A8=
=E9=83=A8=E5=85=91=E6=8D=A2=E4=BA=86=E5=90=A7</p>
</p><hr><p>[=E6=97=B6=E9=80=9F=E4=BA=91]=E8=AF=B7=E5=8F=8A=E6=97=B6=
=E5=9B=9E=E5=A4=8D=E5=B7=A5=E5=8D=95</p><p>2016-09-30 =
11:05:33</p>`: "xxx",
	}
	for k, v := range shouldSuccess {
		header, err := decode(k)
		if err != nil {
			t.Fatalf("decode failed, error:%s\n", err)
		}

		if header != v {
			t.Fatalf("expecting %s, got %s\n", v, header)
		}

		t.Logf("decoded: %s\n", header)
	}
}
