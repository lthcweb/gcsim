package curves

type CharStatCurve int

//nolint:stylecheck,revive // want to match dm here
const (
	GROW_CURVE_HP_S4     CharStatCurve = iota + 1 // "GROW_CURVE_HP_S4"
	GROW_CURVE_ATTACK_S4                          // "GROW_CURVE_ATTACK_S4"
	GROW_CURVE_HP_S5                              // "GROW_CURVE_HP_S5"
	GROW_CURVE_ATTACK_S5                          // "GROW_CURVE_ATTACK_S5"
)

type WeaponStatCurve int

//nolint:stylecheck,revive // want to match dm here
const (
	GROW_CURVE_ATTACK_101 WeaponStatCurve = iota
	GROW_CURVE_ATTACK_102
	GROW_CURVE_ATTACK_103
	GROW_CURVE_ATTACK_104
	GROW_CURVE_ATTACK_105
	GROW_CURVE_CRITICAL_101
	GROW_CURVE_ATTACK_201
	GROW_CURVE_ATTACK_202
	GROW_CURVE_ATTACK_203
	GROW_CURVE_ATTACK_204
	GROW_CURVE_ATTACK_205
	GROW_CURVE_CRITICAL_201
	GROW_CURVE_ATTACK_301
	GROW_CURVE_ATTACK_302
	GROW_CURVE_ATTACK_303
	GROW_CURVE_ATTACK_304
	GROW_CURVE_ATTACK_305
	GROW_CURVE_CRITICAL_301
)
