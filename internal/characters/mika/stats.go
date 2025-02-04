package mika

var (
	attack = [][][]float64{
		//1
		{
			{
				0.4326,
				0.4678,
				0.5031,
				0.5534,
				0.5886,
				0.6288,
				0.6842,
				0.7395,
				0.7948,
				0.8552,
				0.9156,
				0.9759,
				1.0363,
				1.0967,
				1.157,
			},
		},
		//2
		{
			{
				0.415,
				0.4488,
				0.4826,
				0.5308,
				0.5646,
				0.6032,
				0.6563,
				0.7094,
				0.7625,
				0.8204,
				0.8783,
				0.9362,
				0.9941,
				1.052,
				1.1099,
			},
		},
		//3
		{
			{
				0.545,
				0.5894,
				0.6338,
				0.6971,
				0.7415,
				0.7922,
				0.8619,
				0.9316,
				1.0013,
				1.0774,
				1.1534,
				1.2295,
				1.3055,
				1.3816,
				1.4576,
			},
		},
		//4
		{
			{
				0.2761,
				0.2986,
				0.3211,
				0.3532,
				0.3757,
				0.4014,
				0.4367,
				0.472,
				0.5073,
				0.5459,
				0.5844,
				0.6229,
				0.6615,
				0.7,
				0.7385,
			},
			{
				0.2761,
				0.2986,
				0.3211,
				0.3532,
				0.3757,
				0.4014,
				0.4367,
				0.472,
				0.5073,
				0.5459,
				0.5844,
				0.6229,
				0.6615,
				0.7,
				0.7385,
			},
		},
		//5
		{
			{
				0.7087,
				0.7664,
				0.8241,
				0.9065,
				0.9642,
				1.0302,
				1.1208,
				1.2115,
				1.3021,
				1.401,
				1.4999,
				1.5988,
				1.6977,
				1.7966,
				1.8955,
			},
		},
	}

	charge = []float64{
		1.1275,
		1.2192,
		1.311,
		1.4421,
		1.5339,
		1.6387,
		1.783,
		1.9272,
		2.0714,
		2.2287,
		2.386,
		2.5433,
		2.7007,
		2.858,
		3.0153,
	}

	skillPress = []float64{
		0.672,
		0.7224,
		0.7728,
		0.84,
		0.8904,
		0.9408,
		1.008,
		1.0752,
		1.1424,
		1.2096,
		1.2768,
		1.344,
		1.428,
		1.512,
		1.596,
	}

	skillHold = []float64{
		0.84,
		0.903,
		0.966,
		1.05,
		1.113,
		1.176,
		1.26,
		1.344,
		1.428,
		1.512,
		1.596,
		1.68,
		1.785,
		1.89,
		1.995,
	}

	skillExplode = []float64{
		0.252,
		0.2709,
		0.2898,
		0.315,
		0.3339,
		0.3528,
		0.378,
		0.4032,
		0.4284,
		0.4536,
		0.4788,
		0.504,
		0.5355,
		0.567,
		0.5985,
	}

	atkSpdBuff = []float64{
		0.13,
		0.14,
		0.15,
		0.16,
		0.17,
		0.18,
		0.19,
		0.2,
		0.21,
		0.22,
		0.23,
		0.24,
		0.25,
		0.25,
		0.25,
	}

	burstHealFirstP = []float64{
		0.1217,
		0.1308,
		0.1399,
		0.1521,
		0.1612,
		0.1704,
		0.1825,
		0.1947,
		0.2069,
		0.219,
		0.2312,
		0.2434,
		0.2586,
		0.2738,
		0.289,
	}

	burstHealFirstF = []float64{
		1172.0355,
		1289.2554,
		1416.2435,
		1553,
		1699.5248,
		1855.8179,
		2021.8794,
		2197.7092,
		2383.3071,
		2578.6736,
		2783.8083,
		2998.7114,
		3223.3828,
		3457.8225,
		3702.0305,
	}

	burstHealP = []float64{
		0.0243,
		0.0261,
		0.028,
		0.0304,
		0.0322,
		0.034,
		0.0365,
		0.0389,
		0.0413,
		0.0438,
		0.0462,
		0.0486,
		0.0517,
		0.0547,
		0.0578,
	}

	burstHealF = []float64{
		233.9543,
		257.353,
		282.7015,
		310,
		339.2484,
		370.4466,
		403.5947,
		438.6927,
		475.7407,
		514.7385,
		555.6862,
		598.5837,
		643.4312,
		690.2286,
		738.9758,
	}
)
