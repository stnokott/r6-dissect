// Code generated by "stringer -type=Map"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CLUB_HOUSE-837214085]
	_ = x[KAFE_DOSTOYEVSKY-1378191338]
	_ = x[KANAL-1460220617]
	_ = x[YACHT-1767965020]
	_ = x[PRESIDENTIAL_PLANE-2609218856]
	_ = x[CONSULATE-2609221242]
	_ = x[BARTLETT_U-2697268122]
	_ = x[COASTLINE-42090092951]
	_ = x[TOWER-53627213396]
	_ = x[VILLA-88107330328]
	_ = x[FORTRESS-126196841359]
	_ = x[HEREFORD_BASE-127951053400]
	_ = x[THEME_PARK-199824623654]
	_ = x[OREGON-231702797556]
	_ = x[HOUSE-237873412352]
	_ = x[CHALET-259816839773]
	_ = x[SKYSCRAPER-276279025182]
	_ = x[BORDER-305979357167]
	_ = x[FAVELA-329867321446]
	_ = x[BANK-355496559878]
	_ = x[OUTBACK-362605108559]
}

const _Map_name = "CLUB_HOUSEKAFE_DOSTOYEVSKYKANALYACHTPRESIDENTIAL_PLANECONSULATEBARTLETT_UCOASTLINETOWERVILLAFORTRESSHEREFORD_BASETHEME_PARKOREGONHOUSECHALETSKYSCRAPERBORDERFAVELABANKOUTBACK"

var _Map_map = map[Map]string{
	837214085:    _Map_name[0:10],
	1378191338:   _Map_name[10:26],
	1460220617:   _Map_name[26:31],
	1767965020:   _Map_name[31:36],
	2609218856:   _Map_name[36:54],
	2609221242:   _Map_name[54:63],
	2697268122:   _Map_name[63:73],
	42090092951:  _Map_name[73:82],
	53627213396:  _Map_name[82:87],
	88107330328:  _Map_name[87:92],
	126196841359: _Map_name[92:100],
	127951053400: _Map_name[100:113],
	199824623654: _Map_name[113:123],
	231702797556: _Map_name[123:129],
	237873412352: _Map_name[129:134],
	259816839773: _Map_name[134:140],
	276279025182: _Map_name[140:150],
	305979357167: _Map_name[150:156],
	329867321446: _Map_name[156:162],
	355496559878: _Map_name[162:166],
	362605108559: _Map_name[166:173],
}

func (i Map) String() string {
	if str, ok := _Map_map[i]; ok {
		return str
	}
	return "Map(" + strconv.FormatInt(int64(i), 10) + ")"
}
