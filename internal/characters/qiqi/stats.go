package qiqi

var (
	attack = [][][]float64{
		//1
		{
			{
				0.3775,
				0.4083,
				0.439,
				0.4829,
				0.5136,
				0.5488,
				0.597,
				0.6453,
				0.6936,
				0.7463,
				0.799,
				0.8517,
				0.9043,
				0.957,
				1.0097,
			},
		},
		//2
		{
			{
				0.3887,
				0.4204,
				0.452,
				0.4972,
				0.5288,
				0.565,
				0.6147,
				0.6644,
				0.7142,
				0.7684,
				0.8226,
				0.8769,
				0.9311,
				0.9854,
				1.0396,
			},
		},
		// 3 (2 hits)
		{
			{
				0.2417,
				0.2613,
				0.281,
				0.3091,
				0.3288,
				0.3513,
				0.3822,
				0.4131,
				0.444,
				0.4777,
				0.5114,
				0.5451,
				0.5789,
				0.6126,
				0.6463,
			},
			{
				0.2417,
				0.2613,
				0.281,
				0.3091,
				0.3288,
				0.3513,
				0.3822,
				0.4131,
				0.444,
				0.4777,
				0.5114,
				0.5451,
				0.5789,
				0.6126,
				0.6463,
			},
		},
		// 4 (2 hit)
		{
			{
				0.2468,
				0.2669,
				0.287,
				0.3157,
				0.3358,
				0.3588,
				0.3903,
				0.4219,
				0.4535,
				0.4879,
				0.5223,
				0.5568,
				0.5912,
				0.6257,
				0.6601,
			},
			{
				0.2468,
				0.2669,
				0.287,
				0.3157,
				0.3358,
				0.3588,
				0.3903,
				0.4219,
				0.4535,
				0.4879,
				0.5223,
				0.5568,
				0.5912,
				0.6257,
				0.6601,
			},
		},
		//5
		{
			{
				0.6304,
				0.6817,
				0.733,
				0.8063,
				0.8576,
				0.9163,
				0.9969,
				1.0775,
				1.1581,
				1.2461,
				1.3341,
				1.422,
				1.51,
				1.5979,
				1.6859,
			},
		},
	}

	charge = [][]float64{
		{
			0.6433,
			0.6956,
			0.748,
			0.8228,
			0.8752,
			0.935,
			1.0173,
			1.0996,
			1.1818,
			1.2716,
			1.3614,
			1.4511,
			1.5409,
			1.6306,
			1.7204,
		},
		{
			0.6433,
			0.6956,
			0.748,
			0.8228,
			0.8752,
			0.935,
			1.0173,
			1.0996,
			1.1818,
			1.2716,
			1.3614,
			1.4511,
			1.5409,
			1.6306,
			1.7204,
		},
	}

	skillInitialDmg = []float64{
		0.96,
		1.032,
		1.104,
		1.2,
		1.272,
		1.344,
		1.44,
		1.536,
		1.632,
		1.728,
		1.824,
		1.92,
		2.04,
		2.16,
		2.28,
	}

	skillHealOnHitPer = []float64{
		0.1056,
		0.1135,
		0.1214,
		0.132,
		0.1399,
		0.1478,
		0.1584,
		0.169,
		0.1795,
		0.1901,
		0.2006,
		0.2112,
		0.2244,
		0.2376,
		0.2508,
	}

	skillHealOnHitFlat = []float64{
		67,
		74,
		81,
		89,
		98,
		107,
		116,
		126,
		137,
		148,
		160,
		172,
		185,
		199,
		213,
	}

	skillHealContPer = []float64{
		0.696,
		0.7482,
		0.8004,
		0.87,
		0.9222,
		0.9744,
		1.044,
		1.1136,
		1.1832,
		1.2528,
		1.3224,
		1.392,
		1.479,
		1.566,
		1.653,
	}

	skillHealContFlat = []float64{
		451,
		496,
		544,
		597,
		653,
		713,
		777,
		845,
		916,
		991,
		1070,
		1153,
		1239,
		1329,
		1423,
	}

	skillDmgCont = []float64{
		0.36,
		0.387,
		0.414,
		0.45,
		0.477,
		0.504,
		0.54,
		0.576,
		0.612,
		0.648,
		0.684,
		0.72,
		0.765,
		0.81,
		0.855,
	}

	burstDmg = []float64{
		2.848,
		3.0616,
		3.2752,
		3.56,
		3.7736,
		3.9872,
		4.272,
		4.5568,
		4.8416,
		5.1264,
		5.4112,
		5.696,
		6.052,
		6.408,
		6.764,
	}

	burstHealPer = []float64{
		0.9,
		0.9675,
		1.035,
		1.125,
		1.1925,
		1.26,
		1.35,
		1.44,
		1.53,
		1.62,
		1.71,
		1.8,
		1.9125,
		2.025,
		2.1375,
	}

	burstHealFlat = []float64{
		577,
		635,
		698,
		765,
		837,
		914,
		996,
		1083,
		1174,
		1270,
		1371,
		1477,
		1588,
		1703,
		1824,
	}
)
