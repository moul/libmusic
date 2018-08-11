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
	// m2 1 25/24 1.0416666666666667 -25/24 -1.0416666666666667
	// M2 2 9/8 1.125 -9/8 -1.125
	// m3 3 6/5 1.2 -6/5 -1.2
	// M3 4 5/4 1.25 -5/4 -1.25
	// P4 5 4/3 1.3333333333333333 -4/3 -1.3333333333333333
	// d5 6 45/32 1.40625 -45/32 -1.40625
	// P5 7 3/2 1.5 -3/2 -1.5
	// m6 8 8/5 1.6 -8/5 -1.6
	// M6 9 5/3 1.6666666666666667 -5/3 -1.6666666666666667
	// m7 10 9/5 1.8 -9/5 -1.8
	// M7 11 15/8 1.875 -15/8 -1.875
	// P8 12 2/1 2 -2/1 -2
	// m9 13 25/12 2.0833333333333335 -25/12 -2.0833333333333335
	// M9 14 9/4 2.25 -9/4 -2.25
	// m10 15 12/5 2.4 -12/5 -2.4
	// M10 16 5/2 2.5 -5/2 -2.5
	// P11 17 8/3 2.6666666666666665 -8/3 -2.6666666666666665
	// d12 18 45/16 2.8125 -45/16 -2.8125
	// P12 19 3/1 3 -3/1 -3
	// m13 20 16/5 3.2 -16/5 -3.2
	// M13 21 10/3 3.3333333333333335 -10/3 -3.3333333333333335
	// m14 22 18/5 3.6 -18/5 -3.6
	// M14 23 15/4 3.75 -15/4 -3.75
	// P15 24 4/1 4 -4/1 -4
	// 25 25 25/6 4.166666666666667 -25/6 -4.166666666666667
	// 26 26 9/2 4.5 -9/2 -4.5
	// 27 27 24/5 4.8 -24/5 -4.8
	// 28 28 5/1 5 -5/1 -5
	// 29 29 16/3 5.333333333333333 -16/3 -5.333333333333333
	// 30 30 45/8 5.625 -45/8 -5.625
	// 31 31 6/1 6 -6/1 -6
	// 32 32 32/5 6.4 -32/5 -6.4
	// 33 33 20/3 6.666666666666667 -20/3 -6.666666666666667
	// 34 34 36/5 7.2 -36/5 -7.2
	// 35 35 15/2 7.5 -15/2 -7.5
	// 36 36 8/1 8 -8/1 -8
	// 37 37 25/3 8.333333333333334 -25/3 -8.333333333333334
	// 38 38 9/1 9 -9/1 -9
	// 39 39 48/5 9.6 -48/5 -9.6
	// 40 40 10/1 10 -10/1 -10
	// 41 41 32/3 10.666666666666666 -32/3 -10.666666666666666
	// 42 42 45/4 11.25 -45/4 -11.25
	// 43 43 12/1 12 -12/1 -12
	// 44 44 64/5 12.8 -64/5 -12.8
	// 45 45 40/3 13.333333333333334 -40/3 -13.333333333333334
	// 46 46 72/5 14.4 -72/5 -14.4
	// 47 47 15/1 15 -15/1 -15
	// 48 48 16/1 16 -16/1 -16
	// 49 49 50/3 16.666666666666668 -50/3 -16.666666666666668
	// 50 50 18/1 18 -18/1 -18
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
	// m2 1 256/243 1.0534979423868314 -256/243 -1.0534979423868314
	// M2 2 9/8 1.125 -9/8 -1.125
	// m3 3 32/27 1.1851851851851851 -32/27 -1.1851851851851851
	// M3 4 81/64 1.265625 -81/64 -1.265625
	// P4 5 4/3 1.3333333333333333 -4/3 -1.3333333333333333
	// d5 6 729/512 1.423828125 -729/512 -1.423828125
	// P5 7 3/2 1.5 -3/2 -1.5
	// m6 8 128/81 1.5802469135802468 -128/81 -1.5802469135802468
	// M6 9 27/16 1.6875 -27/16 -1.6875
	// m7 10 16/9 1.7777777777777777 -16/9 -1.7777777777777777
	// M7 11 243/128 1.8984375 -243/128 -1.8984375
	// P8 12 2/1 2 -2/1 -2
	// m9 13 512/243 2.1069958847736627 -512/243 -2.1069958847736627
	// M9 14 9/4 2.25 -9/4 -2.25
	// m10 15 64/27 2.3703703703703702 -64/27 -2.3703703703703702
	// M10 16 81/32 2.53125 -81/32 -2.53125
	// P11 17 8/3 2.6666666666666665 -8/3 -2.6666666666666665
	// d12 18 729/256 2.84765625 -729/256 -2.84765625
	// P12 19 3/1 3 -3/1 -3
	// m13 20 256/81 3.1604938271604937 -256/81 -3.1604938271604937
	// M13 21 27/8 3.375 -27/8 -3.375
	// m14 22 32/9 3.5555555555555554 -32/9 -3.5555555555555554
	// M14 23 243/64 3.796875 -243/64 -3.796875
	// P15 24 4/1 4 -4/1 -4
	// 25 25 1024/243 4.2139917695473255 -1024/243 -4.2139917695473255
	// 26 26 9/2 4.5 -9/2 -4.5
	// 27 27 128/27 4.7407407407407405 -128/27 -4.7407407407407405
	// 28 28 81/16 5.0625 -81/16 -5.0625
	// 29 29 16/3 5.333333333333333 -16/3 -5.333333333333333
	// 30 30 729/128 5.6953125 -729/128 -5.6953125
	// 31 31 6/1 6 -6/1 -6
	// 32 32 512/81 6.320987654320987 -512/81 -6.320987654320987
	// 33 33 27/4 6.75 -27/4 -6.75
	// 34 34 64/9 7.111111111111111 -64/9 -7.111111111111111
	// 35 35 243/32 7.59375 -243/32 -7.59375
	// 36 36 8/1 8 -8/1 -8
	// 37 37 2048/243 8.427983539094651 -2048/243 -8.427983539094651
	// 38 38 9/1 9 -9/1 -9
	// 39 39 256/27 9.481481481481481 -256/27 -9.481481481481481
	// 40 40 81/8 10.125 -81/8 -10.125
	// 41 41 32/3 10.666666666666666 -32/3 -10.666666666666666
	// 42 42 729/64 11.390625 -729/64 -11.390625
	// 43 43 12/1 12 -12/1 -12
	// 44 44 1024/81 12.641975308641975 -1024/81 -12.641975308641975
	// 45 45 27/2 13.5 -27/2 -13.5
	// 46 46 128/9 14.222222222222221 -128/9 -14.222222222222221
	// 47 47 243/16 15.1875 -243/16 -15.1875
	// 48 48 16/1 16 -16/1 -16
	// 49 49 4096/243 16.855967078189302 -4096/243 -16.855967078189302
	// 50 50 18/1 18 -18/1 -18
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
