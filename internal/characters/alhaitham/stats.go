package alhaitham

var (
	rushAtk = []float64{
		193.6 / 100.0,
		208.12 / 100.0,
		222.64 / 100.0,
		242 / 100.0,
		256.52 / 100.0,
		271.04 / 100.0,
		290.4 / 100.0,
		309.76 / 100.0,
		329.12 / 100.0,
		348.48 / 100.0,
		367.84 / 100.0,
		387.2 / 100.0,
		411.4 / 100.0,
		435.6 / 100.0,
		459.8 / 100.0,
	}
	rushEm = []float64{
		154.88 / 100.0,
		166.5 / 100.0,
		178.11 / 100.0,
		193.6 / 100.0,
		205.22 / 100.0,
		216.83 / 100.0,
		232.32 / 100.0,
		247.81 / 100.0,
		263.3 / 100.0,
		278.78 / 100.0,
		294.27 / 100.0,
		309.76 / 100.0,
		329.12 / 100.0,
		348.48 / 100.0,
		367.84 / 100.0,
	}

	mirrorAtk = []float64{
		67.2 / 100.0,
		72.24 / 100.0,
		77.28 / 100.0,
		84 / 100.0,
		89.04 / 100.0,
		94.08 / 100.0,
		100.8 / 100.0,
		107.52 / 100.0,
		114.24 / 100.0,
		120.96 / 100.0,
		127.68 / 100.0,
		134.4 / 100.0,
		142.8 / 100.0,
		151.2 / 100.0,
		159.6 / 100.0,
	}

	mirrorEm = []float64{
		134.4 / 100.0,
		144.48 / 100.0,
		154.56 / 100.0,
		168 / 100.0,
		178.08 / 100.0,
		188.16 / 100.0,
		201.6 / 100.0,
		215.04 / 100.0,
		228.48 / 100.0,
		241.92 / 100.0,
		255.36 / 100.0,
		268.8 / 100.0,
		285.6 / 100.0,
		302.4 / 100.0,
		319.2 / 100.0,
	}

	burstAtk = []float64{
		121.6 / 100,
		130.72 / 100,
		139.84 / 100,
		152 / 100,
		161.12 / 100,
		170.24 / 100,
		182.4 / 100,
		194.56 / 100,
		206.72 / 100,
		218.88 / 100,
		231.04 / 100,
		243.2 / 100,
		258.4 / 100,
		273.6 / 100,
		288.8 / 100,
	}
	burstEm = []float64{
		97.28 / 100,
		104.58 / 100,
		111.87 / 100,
		121.6 / 100,
		128.9 / 100,
		136.19 / 100,
		145.92 / 100,
		155.65 / 100,
		165.38 / 100,
		175.1 / 100,
		184.83 / 100,
		194.56 / 100,
		206.72 / 100,
		218.88 / 100,
		231.04 / 100,
	}

	attack = [][][]float64{
		{ //1
			{
				49.53 / 100.0,
				53.56 / 100.0,
				57.59 / 100.0,
				63.35 / 100.0,
				67.38 / 100.0,
				71.99 / 100.0,
				78.32 / 100.0,
				84.65 / 100.0,
				90.99 / 100.0,
				97.9 / 100.0,
				104.81 / 100.0,
				111.72 / 100.0,
				118.63 / 100.0,
				125.54 / 100.0,
				132.45 / 100.0,
			},
		},
		{ //2
			{
				50.75 / 100.0,
				54.88 / 100.0,
				59.01 / 100.0,
				64.91 / 100.0,
				69.04 / 100.0,
				73.76 / 100.0,
				80.26 / 100.0,
				86.75 / 100.0,
				93.24 / 100.0,
				100.32 / 100.0,
				107.4 / 100.0,
				114.48 / 100.0,
				121.56 / 100.0,
				128.64 / 100.0,
				135.73 / 100.0,
			},
		},
		{ //3
			{
				34.18 / 100.0,
				36.96 / 100.0,
				39.74 / 100.0,
				43.72 / 100.0,
				46.5 / 100.0,
				49.68 / 100.0,
				54.05 / 100.0,
				58.42 / 100.0,
				62.79 / 100.0,
				67.56 / 100.0,
				72.33 / 100.0,
				77.1 / 100.0,
				81.87 / 100.0,
				86.64 / 100.0,
				91.41 / 100.0,
			},
			{
				34.18 / 100.0,
				36.96 / 100.0,
				39.74 / 100.0,
				43.72 / 100.0,
				46.5 / 100.0,
				49.68 / 100.0,
				54.05 / 100.0,
				58.42 / 100.0,
				62.79 / 100.0,
				67.56 / 100.0,
				72.33 / 100.0,
				77.1 / 100.0,
				81.87 / 100.0,
				86.64 / 100.0,
				91.41 / 100.0,
			},
		},
		{ //4
			{
				66.77 / 100.0,
				72.2 / 100.0,
				77.64 / 100.0,
				85.4 / 100.0,
				90.84 / 100.0,
				97.05 / 100.0,
				105.59 / 100.0,
				114.13 / 100.0,
				122.67 / 100.0,
				131.98 / 100.0,
				141.3 / 100.0,
				150.62 / 100.0,
				159.93 / 100.0,
				169.25 / 100.0,
				178.57 / 100.0,
			},
		},
		{ //5
			{
				83.85 / 100.0,
				90.68 / 100.0,
				97.5 / 100.0,
				107.25 / 100.0,
				114.08 / 100.0,
				121.88 / 100.0,
				132.6 / 100.0,
				143.33 / 100.0,
				154.05 / 100.0,
				165.75 / 100.0,
				177.45 / 100.0,
				189.15 / 100.0,
				200.85 / 100.0,
				212.55 / 100.0,
				224.25 / 100.0,
			},
		},
	}
	charge = [][]float64{
		{
			55.26 / 100.0,
			59.75 / 100.0,
			64.25 / 100.0,
			70.67 / 100.0,
			75.17 / 100.0,
			80.31 / 100.0,
			87.38 / 100.0,
			94.45 / 100.0,
			101.51 / 100.0,
			109.22 / 100.0,
			116.94 / 100.0,
			124.64 / 100.0,
			132.35 / 100.0,
			140.07 / 100.0,
			147.77 / 100.0,
		},
		{
			55.26 / 100.0,
			59.75 / 100.0,
			64.25 / 100.0,
			70.67 / 100.0,
			75.17 / 100.0,
			80.31 / 100.0,
			87.38 / 100.0,
			94.45 / 100.0,
			101.51 / 100.0,
			109.22 / 100.0,
			116.94 / 100.0,
			124.64 / 100.0,
			132.35 / 100.0,
			140.07 / 100.0,
			147.77 / 100.0,
		},
	}
	lowPlunge = []float64{
		127.84 / 100.0,
		138.24 / 100.0,
		148.65 / 100.0,
		163.51 / 100.0,
		173.92 / 100.0,
		185.81 / 100.0,
		202.16 / 100.0,
		218.51 / 100.0,
		234.86 / 100.0,
		252.7 / 100.0,
		270.54 / 100.0,
		288.38 / 100.0,
		306.22 / 100.0,
		324.05 / 100.0,
		341.89 / 100.0,
	}
)
