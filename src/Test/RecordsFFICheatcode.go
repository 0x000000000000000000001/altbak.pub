package Test_RecordsFFICheatcode


type player_cheatcode struct {
	x, y, vx, vy int
}

func RunRecordsFFICheatcode(limit int) int {
	p := player_cheatcode{x: 0, y: 0, vx: 1, vy: 1}
	for i := 0; i < int(limit); i++ {
		p.x = p.x + p.vx
		p.y = p.y + p.vy
	}
	return (p.x + p.y)
}

