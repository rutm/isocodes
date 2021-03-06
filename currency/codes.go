package currency

var currencies = []*currency{
	{"AED", "UAE Dirham", 784},
	{"AFN", "Afghani", 971},
	{"ALL", "Lek", 8},
	{"AMD", "Armenian Dram", 51},
	{"ANG", "Netherlands Antillian Guilder", 532},
	{"AOA", "Kwanza", 973},
	{"ARS", "Argentine Peso", 32},
	{"AUD", "Australian Dollar", 36},
	{"AWG", "Aruban Guilder", 533},
	{"AZN", "Azerbaijanian Manat", 944},
	{"BAM", "Convertible Marks", 977},
	{"BBD", "Barbados Dollar", 52},
	{"BDT", "Taka", 50},
	{"BGN", "Bulgarian Lev", 975},
	{"BHD", "Bahraini Dinar", 48},
	{"BIF", "Burundi Franc", 108},
	{"BMD", "Bermudian Dollar", 60},
	{"BND", "Brunei Dollar", 96},
	{"BOB", "Boliviano", 68},
	{"BOV", "Mvdol", 984},
	{"BRL", "Brazilian Real", 986},
	{"BSD", "Bahamian Dollar", 44},
	{"BTN", "Ngultrum", 64},
	{"BWP", "Pula", 72},
	{"BYR", "Belarussian Ruble", 974},
	{"BZD", "Belize Dollar", 84},
	{"CAD", "Canadian Dollar", 124},
	{"CDF", "Congolese Franc", 976},
	{"CHE", "WIR Euro", 947},
	{"CHF", "Swiss Franc", 756},
	{"CHW", "WIR Franc", 948},
	{"CLF", "Unidades de fomento", 990},
	{"CLP", "Chilean Peso", 152},
	{"CNY", "Yuan Renminbi", 156},
	{"COP", "Colombian Peso", 170},
	{"COU", "Unidad de Valor Real", 970},
	{"CRC", "Costa Rican Colon", 188},
	{"CUP", "Cuban Peso", 192},
	{"CVE", "Cape Verde Escudo", 132},
	{"CYP", "Cyprus Pound", 196},
	{"CZK", "Czech Koruna", 203},
	{"DJF", "Djibouti Franc", 262},
	{"DKK", "Danish Krone", 208},
	{"DOP", "Dominican Peso", 214},
	{"DZD", "Algerian Dinar", 12},
	{"EEK", "Kroon", 233},
	{"EGP", "Egyptian Pound", 818},
	{"ERN", "Nakfa", 232},
	{"ETB", "Ethiopian Birr", 230},
	{"EUR", "Euro", 978},
	{"FJD", "Fiji Dollar", 242},
	{"FKP", "Falkland Islands Pound", 238},
	{"GBP", "Pound Sterling", 826},
	{"GEL", "Lari", 981},
	{"GHS", "Ghana Cedi", 936},
	{"GIP", "Gibraltar Pound", 292},
	{"GMD", "Dalasi", 270},
	{"GNF", "Guinea Franc", 324},
	{"GTQ", "Quetzal", 320},
	{"GYD", "Guyana Dollar", 328},
	{"HKD", "Hong Kong Dollar", 344},
	{"HNL", "Lempira", 340},
	{"HRK", "Croatian Kuna", 191},
	{"HTG", "Gourde", 332},
	{"HUF", "Forint", 348},
	{"IDR", "Rupiah", 360},
	{"ILS", "New Israeli Sheqel", 376},
	{"INR", "Indian Rupee", 356},
	{"IQD", "Iraqi Dinar", 368},
	{"IRR", "Iranian Rial", 364},
	{"ISK", "Iceland Krona", 352},
	{"JMD", "Jamaican Dollar", 388},
	{"JOD", "Jordanian Dinar", 400},
	{"JPY", "Yen", 392},
	{"KES", "Kenyan Shilling", 404},
	{"KGS", "Som", 417},
	{"KHR", "Riel", 116},
	{"KMF", "Comoro Franc", 174},
	{"KPW", "North Korean Won", 408},
	{"KRW", "Won", 410},
	{"KWD", "Kuwaiti Dinar", 414},
	{"KYD", "Cayman Islands Dollar", 136},
	{"KZT", "Tenge", 398},
	{"LAK", "Kip", 418},
	{"LBP", "Lebanese Pound", 422},
	{"LKR", "Sri Lanka Rupee", 144},
	{"LRD", "Liberian Dollar", 430},
	{"LSL", "Loti", 426},
	{"LTL", "Lithuanian Litas", 440},
	{"LVL", "Latvian Lats", 428},
	{"LYD", "Libyan Dinar", 434},
	{"MAD", "Moroccan Dirham", 504},
	{"MDL", "Moldovan Leu", 498},
	{"MGA", "Malagasy Ariary", 969},
	{"MKD", "Denar", 807},
	{"MMK", "Kyat", 104},
	{"MNT", "Tugrik", 496},
	{"MOP", "Pataca", 446},
	{"MRO", "Ouguiya", 478},
	{"MTL", "Maltese Lira", 470},
	{"MUR", "Mauritius Rupee", 480},
	{"MVR", "Rufiyaa", 462},
	{"MWK", "Kwacha", 454},
	{"MXN", "Mexican Peso", 484},
	{"MXV", "Mexican Unidad de Inversion (UDI)", 979},
	{"MYR", "Malaysian Ringgit", 458},
	{"MZN", "Metical", 943},
	{"NAD", "Namibia Dollar", 516},
	{"NGN", "Naira", 566},
	{"NIO", "Cordoba Oro", 558},
	{"NOK", "Norwegian Krone", 578},
	{"NPR", "Nepalese Rupee", 524},
	{"NZD", "New Zealand Dollar", 554},
	{"OMR", "Rial Omani", 512},
	{"PAB", "Balboa", 590},
	{"PEN", "Nuevo Sol", 604},
	{"PGK", "Kina", 598},
	{"PHP", "Philippine Peso", 608},
	{"PKR", "Pakistan Rupee", 586},
	{"PLN", "Zloty", 985},
	{"PYG", "Guarani", 600},
	{"QAR", "Qatari Rial", 634},
	{"RON", "New Leu", 946},
	{"RSD", "Serbian Dinar", 941},
	{"RUB", "Russian Ruble", 643},
	{"RWF", "Rwanda Franc", 646},
	{"SAR", "Saudi Riyal", 682},
	{"SBD", "Solomon Islands Dollar", 90},
	{"SCR", "Seychelles Rupee", 690},
	{"SDG", "Sudanese Pound", 938},
	{"SEK", "Swedish Krona", 752},
	{"SGD", "Singapore Dollar", 702},
	{"SHP", "Saint Helena Pound", 654},
	{"SLL", "Leone", 694},
	{"SOS", "Somali Shilling", 706},
	{"SRD", "Surinam Dollar", 968},
	{"STD", "Dobra", 678},
	{"SVC", "El Salvador Colon", 222},
	{"SYP", "Syrian Pound", 760},
	{"SZL", "Lilangeni", 748},
	{"THB", "Baht", 764},
	{"TJS", "Somoni", 972},
	{"TMM", "Manat", 795},
	{"TND", "Tunisian Dinar", 788},
	{"TOP", "Pa'anga", 776},
	{"TRY", "New Turkish Lira", 949},
	{"TTD", "Trinidad and Tobago Dollar", 780},
	{"TWD", "New Taiwan Dollar", 901},
	{"TZS", "Tanzanian Shilling", 834},
	{"UAH", "Hryvnia", 980},
	{"UGX", "Uganda Shilling", 800},
	{"USD", "US Dollar", 840},
	{"USN", "US Dollar (Next day)", 997},
	{"USS", "US Dollar (Same day)", 998},
	{"UYI", "Uruguay Peso en Unidades Indexadas", 940},
	{"UYU", "Peso Uruguayo", 858},
	{"UZS", "Uzbekistan Sum", 860},
	{"VEF", "Bolivar Fuerte", 937},
	{"VND", "Dong", 704},
	{"VUV", "Vatu", 548},
	{"WST", "Tala", 882},
	{"XAF", "CFA Franc BEAC", 950},
	{"XAG", "Silver", 961},
	{"XAU", "Gold", 959},
	{"XBA", "European Composite Unit (EURCO)", 955},
	{"XBB", "European Monetary Unit (E.M.U.-6)", 956},
	{"XBC", "European Unit of Account 9 (E.U.A.-9)", 957},
	{"XBD", "European Unit of Account 17 (E.U.A.-17)", 958},
	{"XCD", "East Caribbean Dollar", 951},
	{"XDR", "Special Drawing Rights", 960},
	{"XFO", "Gold-Franc", 0},
	{"XFU", "UIC-Franc", 0},
	{"XOF", "CFA Franc BCEAO", 952},
	{"XPD", "Palladium", 964},
	{"XPF", "CFP Franc", 953},
	{"XPT", "Platinum", 962},
	{"XTS", "Code for testing purposes", 963},
	{"XXX", "No currency", 999},
	{"YER", "Yemeni Rial", 886},
	{"ZAR", "Rand", 710},
	{"ZMK", "Zambian Kwacha", 894},
	{"ZWD", "Zimbabwe Dollar", 716},
}
