package leetcode

// 在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
//
// 每个矩形由其左下顶点和右上顶点坐标表示，如图所示。
//
//
//
// 示例:
//
// 输入: -3, 0, 3, 4, 0, -1, 9, 2
// 输出: 45

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sRecABCD := (C - A) * (D - B)
	sRecEFGH := (G - E) * (H - F)

	sMiddle := 0

	// 得到重叠面积的边界x, y值
	lx := max(E, A)
	rx := min(C, G)
	dy := max(F, B)
	uy := min(H, D)

	if E > C || F > D || G < A || H < B {
		return sRecABCD + sRecEFGH
	}

	sMiddle = (lx - rx) * (dy - uy)

	return sRecABCD + sRecEFGH - sMiddle
}
