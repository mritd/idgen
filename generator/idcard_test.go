package generator

import "testing"

func TestGetIssueOrg(t *testing.T) {
	t.Log(GetIssueOrg())
}

func TestGetValidPeriod(t *testing.T) {
	t.Log(GetValidPeriod())
}

func TestGetIDNo(t *testing.T) {
	t.Log(GetIDNo())
}

func TestVerifyCode(t *testing.T) {
	t.Log(VerifyCode("636706198006242277"))
}

func TestRandDate(t *testing.T) {
	t.Log(RandDate())
}

func BenchmarkGetIssueOrg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetIssueOrg())
	}
}

func BenchmarkGetValidPeriod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetValidPeriod())
	}
}

func BenchmarkGetIDNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetIDNo())
	}
}

func BenchmarkVerifyCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(VerifyCode("636706198006242277"))
	}
}
