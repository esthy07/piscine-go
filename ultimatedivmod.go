package piscine

func UltimateDivMod(a *int, b *int) {
	x := *div
	*div = *div / *mod
	*mod = x % *mod
}