package libmusic

import "fmt"

func ExampleNote_Frequency() {
	// low-frequency notes
	fmt.Println(C0, C0.Frequency())
	fmt.Println(D0, D0.Frequency())
	fmt.Println(E0, E0.Frequency())
	fmt.Println(F0, F0.Frequency())
	fmt.Println(G0, G0.Frequency())
	fmt.Println(A0, A0.Frequency())
	fmt.Println(B0, B0.Frequency())

	// octaves
	fmt.Println(A0, A0.Frequency())
	fmt.Println(A1, A1.Frequency())
	fmt.Println(A2, A2.Frequency())
	fmt.Println(A3, A3.Frequency())
	fmt.Println(A4, A4.Frequency())
	fmt.Println(A5, A5.Frequency())
	fmt.Println(A6, A6.Frequency())
	fmt.Println(A7, A7.Frequency())
	fmt.Println(A8, A8.Frequency())

	// Output:
	// C0 16.352㎐
	// D0 18.354㎐
	// E0 20.602㎐
	// F0 21.827㎐
	// G0 24.5㎐
	// A0 27.5㎐
	// B0 30.868㎐
	// A0 27.5㎐
	// A1 55㎐
	// A2 110㎐
	// A3 220㎐
	// A4 440㎐
	// A5 880㎐
	// A6 1760㎐
	// A7 3520㎐
	// A8 7040㎐
}

func ExampleNote_JustScaleFrequency() {
	// low-frequency notes
	fmt.Println(C0, C0.JustScaleFrequency())
	fmt.Println(D0, D0.JustScaleFrequency())
	fmt.Println(E0, E0.JustScaleFrequency())
	fmt.Println(F0, F0.JustScaleFrequency())
	fmt.Println(G0, G0.JustScaleFrequency())
	fmt.Println(A0, A0.JustScaleFrequency())
	fmt.Println(B0, B0.JustScaleFrequency())

	// octaves
	fmt.Println(A0, A0.JustScaleFrequency())
	fmt.Println(A1, A1.JustScaleFrequency())
	fmt.Println(A2, A2.JustScaleFrequency())
	fmt.Println(A3, A3.JustScaleFrequency())
	fmt.Println(A4, A4.JustScaleFrequency())
	fmt.Println(A5, A5.JustScaleFrequency())
	fmt.Println(A6, A6.JustScaleFrequency())
	fmt.Println(A7, A7.JustScaleFrequency())
	fmt.Println(A8, A8.JustScaleFrequency())

	// Output:
	// C0 16.5㎐
	// D0 18.333㎐
	// E0 20.625㎐
	// F0 22㎐
	// G0 24.444㎐
	// A0 27.5㎐
	// B0 30.556㎐
	// A0 27.5㎐
	// A1 55㎐
	// A2 110㎐
	// A3 220㎐
	// A4 440㎐
	// A5 880㎐
	// A6 1760㎐
	// A7 3520㎐
	// A8 7040㎐
}

func ExampleHz_WaveLength() {
	// low-frequency notes
	fmt.Println(C0, C0.Frequency(), C0.Frequency().WaveLength())
	fmt.Println(D0, D0.Frequency(), D0.Frequency().WaveLength())
	fmt.Println(E0, E0.Frequency(), E0.Frequency().WaveLength())
	fmt.Println(F0, F0.Frequency(), F0.Frequency().WaveLength())
	fmt.Println(G0, G0.Frequency(), G0.Frequency().WaveLength())
	fmt.Println(A0, A0.Frequency(), A0.Frequency().WaveLength())
	fmt.Println(B0, B0.Frequency(), B0.Frequency().WaveLength())

	// octaves
	fmt.Println(A0, A0.Frequency(), A0.Frequency().WaveLength())
	fmt.Println(A1, A1.Frequency(), A1.Frequency().WaveLength())
	fmt.Println(A2, A2.Frequency(), A2.Frequency().WaveLength())
	fmt.Println(A3, A3.Frequency(), A3.Frequency().WaveLength())
	fmt.Println(A4, A4.Frequency(), A4.Frequency().WaveLength())
	fmt.Println(A5, A5.Frequency(), A5.Frequency().WaveLength())
	fmt.Println(A6, A6.Frequency(), A6.Frequency().WaveLength())
	fmt.Println(A7, A7.Frequency(), A7.Frequency().WaveLength())
	fmt.Println(A8, A8.Frequency(), A8.Frequency().WaveLength())

	// Output:
	// C0 16.352㎐ 21.099m/s
	// D0 18.354㎐ 18.797m/s
	// E0 20.602㎐ 16.746m/s
	// F0 21.827㎐ 15.806m/s
	// G0 24.5㎐ 14.082m/s
	// A0 27.5㎐ 12.545m/s
	// B0 30.868㎐ 11.177m/s
	// A0 27.5㎐ 12.545m/s
	// A1 55㎐ 6.273m/s
	// A2 110㎐ 3.136m/s
	// A3 220㎐ 1.568m/s
	// A4 440㎐ 0.784m/s
	// A5 880㎐ 0.392m/s
	// A6 1760㎐ 0.196m/s
	// A7 3520㎐ 0.098m/s
	// A8 7040㎐ 0.049m/s
}

func ExampleInterval_JustScaleRatio() {
	intervals := []Interval{}
	for i := 0; i <= 50; i++ {
		intervals = append(intervals, Interval(i))
	}
	for _, interval := range intervals {
		ratio := interval.JustScaleRatio()
		negRatio := (-interval).JustScaleRatio()
		fmt.Println(interval, interval.Int(), ratio, ratio.Float(), negRatio, negRatio.Float())
	}
	// Output:
	// P1 0 1/1 1 1/1 1
	// m2 1 25/24 1.0416666666666667 24/25 0.96
	// M2 2 9/8 1.125 8/9 0.8888888888888888
	// m3 3 6/5 1.2 5/6 0.8333333333333334
	// M3 4 5/4 1.25 4/5 0.8
	// P4 5 4/3 1.3333333333333333 3/4 0.75
	// d5 6 45/32 1.40625 32/45 0.7111111111111111
	// P5 7 3/2 1.5 2/3 0.6666666666666666
	// m6 8 8/5 1.6 5/8 0.625
	// M6 9 5/3 1.6666666666666667 3/5 0.6
	// m7 10 9/5 1.8 5/9 0.5555555555555556
	// M7 11 15/8 1.875 8/15 0.5333333333333333
	// P8 12 2/1 2 1/2 0.5
	// m9 13 25/12 2.0833333333333335 12/25 0.48
	// M9 14 9/4 2.25 4/9 0.4444444444444444
	// m10 15 12/5 2.4 5/12 0.4166666666666667
	// M10 16 5/2 2.5 2/5 0.4
	// P11 17 8/3 2.6666666666666665 3/8 0.375
	// d12 18 45/16 2.8125 16/45 0.35555555555555557
	// P12 19 3/1 3 1/3 0.3333333333333333
	// m13 20 16/5 3.2 5/16 0.3125
	// M13 21 10/3 3.3333333333333335 3/10 0.3
	// m14 22 18/5 3.6 5/18 0.2777777777777778
	// M14 23 15/4 3.75 4/15 0.26666666666666666
	// P15 24 4/1 4 1/4 0.25
	// 25 25 25/6 4.166666666666667 6/25 0.24
	// 26 26 9/2 4.5 2/9 0.2222222222222222
	// 27 27 24/5 4.8 5/24 0.20833333333333334
	// 28 28 5/1 5 1/5 0.2
	// 29 29 16/3 5.333333333333333 3/16 0.1875
	// 30 30 45/8 5.625 8/45 0.17777777777777778
	// 31 31 6/1 6 1/6 0.16666666666666666
	// 32 32 32/5 6.4 5/32 0.15625
	// 33 33 20/3 6.666666666666667 3/20 0.15
	// 34 34 36/5 7.2 5/36 0.1388888888888889
	// 35 35 15/2 7.5 2/15 0.13333333333333333
	// 36 36 8/1 8 1/8 0.125
	// 37 37 25/3 8.333333333333334 3/25 0.12
	// 38 38 9/1 9 1/9 0.1111111111111111
	// 39 39 48/5 9.6 5/48 0.10416666666666667
	// 40 40 10/1 10 1/10 0.1
	// 41 41 32/3 10.666666666666666 3/32 0.09375
	// 42 42 45/4 11.25 4/45 0.08888888888888889
	// 43 43 12/1 12 1/12 0.08333333333333333
	// 44 44 64/5 12.8 5/64 0.078125
	// 45 45 40/3 13.333333333333334 3/40 0.075
	// 46 46 72/5 14.4 5/72 0.06944444444444445
	// 47 47 15/1 15 1/15 0.06666666666666667
	// 48 48 16/1 16 1/16 0.0625
	// 49 49 50/3 16.666666666666668 3/50 0.06
	// 50 50 18/1 18 1/18 0.05555555555555555
}

func ExampleInterval_PythagoreanRatio() {
	intervals := []Interval{}
	for i := 0; i <= 50; i++ {
		intervals = append(intervals, Interval(i))
	}
	for _, interval := range intervals {
		ratio := interval.PythagoreanRatio()
		negRatio := (-interval).PythagoreanRatio()
		fmt.Println(interval, interval.Int(), ratio, ratio.Float(), negRatio, negRatio.Float())
	}
	// Output:
	// P1 0 1/1 1 1/1 1
	// m2 1 256/243 1.0534979423868314 243/256 0.94921875
	// M2 2 9/8 1.125 8/9 0.8888888888888888
	// m3 3 32/27 1.1851851851851851 27/32 0.84375
	// M3 4 81/64 1.265625 64/81 0.7901234567901234
	// P4 5 4/3 1.3333333333333333 3/4 0.75
	// d5 6 729/512 1.423828125 512/729 0.7023319615912208
	// P5 7 3/2 1.5 2/3 0.6666666666666666
	// m6 8 128/81 1.5802469135802468 81/128 0.6328125
	// M6 9 27/16 1.6875 16/27 0.5925925925925926
	// m7 10 16/9 1.7777777777777777 9/16 0.5625
	// M7 11 243/128 1.8984375 128/243 0.5267489711934157
	// P8 12 2/1 2 1/2 0.5
	// m9 13 512/243 2.1069958847736627 243/512 0.474609375
	// M9 14 9/4 2.25 4/9 0.4444444444444444
	// m10 15 64/27 2.3703703703703702 27/64 0.421875
	// M10 16 81/32 2.53125 32/81 0.3950617283950617
	// P11 17 8/3 2.6666666666666665 3/8 0.375
	// d12 18 729/256 2.84765625 256/729 0.3511659807956104
	// P12 19 3/1 3 1/3 0.3333333333333333
	// m13 20 256/81 3.1604938271604937 81/256 0.31640625
	// M13 21 27/8 3.375 8/27 0.2962962962962963
	// m14 22 32/9 3.5555555555555554 9/32 0.28125
	// M14 23 243/64 3.796875 64/243 0.26337448559670784
	// P15 24 4/1 4 1/4 0.25
	// 25 25 1024/243 4.2139917695473255 243/1024 0.2373046875
	// 26 26 9/2 4.5 2/9 0.2222222222222222
	// 27 27 128/27 4.7407407407407405 27/128 0.2109375
	// 28 28 81/16 5.0625 16/81 0.19753086419753085
	// 29 29 16/3 5.333333333333333 3/16 0.1875
	// 30 30 729/128 5.6953125 128/729 0.1755829903978052
	// 31 31 6/1 6 1/6 0.16666666666666666
	// 32 32 512/81 6.320987654320987 81/512 0.158203125
	// 33 33 27/4 6.75 4/27 0.14814814814814814
	// 34 34 64/9 7.111111111111111 9/64 0.140625
	// 35 35 243/32 7.59375 32/243 0.13168724279835392
	// 36 36 8/1 8 1/8 0.125
	// 37 37 2048/243 8.427983539094651 243/2048 0.11865234375
	// 38 38 9/1 9 1/9 0.1111111111111111
	// 39 39 256/27 9.481481481481481 27/256 0.10546875
	// 40 40 81/8 10.125 8/81 0.09876543209876543
	// 41 41 32/3 10.666666666666666 3/32 0.09375
	// 42 42 729/64 11.390625 64/729 0.0877914951989026
	// 43 43 12/1 12 1/12 0.08333333333333333
	// 44 44 1024/81 12.641975308641975 81/1024 0.0791015625
	// 45 45 27/2 13.5 2/27 0.07407407407407407
	// 46 46 128/9 14.222222222222221 9/128 0.0703125
	// 47 47 243/16 15.1875 16/243 0.06584362139917696
	// 48 48 16/1 16 1/16 0.0625
	// 49 49 4096/243 16.855967078189302 243/4096 0.059326171875
	// 50 50 18/1 18 1/18 0.05555555555555555
}

func ExampleNote_ByFrequency() {
	targets := []Hz{440, 880, 1, 666, 42, 4242.4242, 1000, 1234}
	for _, target := range targets {
		note := NoteByFrequency(target)
		freq := note.Frequency()
		diff := Hz(target - freq).Relative()
		fmt.Printf("target=%s, note=%s, note.freq=%s, diff=%s\n", freq.String(), note.String(), freq, diff)
	}
	// Output:
	// target=440㎐, note=A4, note.freq=440㎐, diff=0㎐
	// target=880㎐, note=A5, note.freq=880㎐, diff=0㎐
	// target=1.022㎐, note=C-4, note.freq=1.022㎐, diff=0.022㎐
	// target=659.255㎐, note=E5, note.freq=659.255㎐, diff=6.745㎐
	// target=41.203㎐, note=E1, note.freq=41.203㎐, diff=0.797㎐
	// target=4186.009㎐, note=C8, note.freq=4186.009㎐, diff=56.415㎐
	// target=987.767㎐, note=B5, note.freq=987.767㎐, diff=12.233㎐
	// target=1244.508㎐, note=D♯6, note.freq=1244.508㎐, diff=10.508㎐
}
