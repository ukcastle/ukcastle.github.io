package pop

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

}
func PopCountLoop(x uint64) int {
	// 단일 표현식 대신 루프 사용하고 성능 비교

	cnt := 0

	for i := 0; i < 8; i++ {
		cnt += int(pc[byte(x>>(i*8))])
	}

	return cnt
}

func PopCount64Bit(x uint64) int {
	// 인수를 시프트시키면서 제일 오른쪽 비트를 매번 테스트해 비트 수를 세는 함수
	return 0
}

func PopCountRemoveRight(x uint64) int {
	//표현식 x&(x-1)은 제일 오른쪽의 0이 아닌 비트를 지움, 이를 이용해보자
	return 0
}
