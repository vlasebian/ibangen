package main

type bank struct {
	name        string
	bic         string
	code        string
	branchCodes []string
}

var dutchBanks = []bank{
	{name: "ING Group", bic: "INGBNL2A", code: "INGB"},
	{name: "ABN Amro", bic: "ABNANL2A", code: "ABNA"},
	{name: "Rabobank", bic: "RABONL2U", code: "RABO"},
	{name: "SNS Bank", bic: "SNSBNL2A", code: "SNSB"},
	{name: "ASN Bank", bic: "ASNBNL21", code: "ASNB"},
	{name: "Triodos Bank", bic: "TRIONL2U", code: "TRIO"},
	{name: "Van Lanschot", bic: "FVLBNL22", code: "FVLB"},
	{name: "NIBC Bank", bic: "NIBCNL2A", code: "NIBB"},
	{name: "RegioBank", bic: "RBRBNL21", code: "RBRB"},
	{name: "Knab", bic: "KNABNL2H", code: "KNAB"},
}

var germanBanks = []bank{
	{name: "Deutsche Bank", bic: "DEUTDEFF", code: "10070000"},
	{name: "Commerzbank", bic: "COBADEFF", code: "10040000"},
	{name: "KfW Bank", bic: "KFWIDEFF", code: "50020000"},
	{name: "DZ Bank", bic: "GENODEFF", code: "50060400"},
	{name: "UniCredit Bank", bic: "HYVEDEMM", code: "70020270"},
	{name: "Landesbank Baden-Württemberg (LBBW)", bic: "SOLADEST", code: "60050101"},
	{name: "Landesbank Hessen-Thüringen (Helaba)", bic: "HELADEFF", code: "50050000"},
	{name: "Norddeutsche Landesbank (NordLB)", bic: "NOLADE2HXXX", code: "25050000"},
	{name: "Landesbank Berlin (LBB)", bic: "BELADEBE", code: "10050000"},
	{name: "Bayerische Landesbank (BayernLB)", bic: "BYLADEMM", code: "70050000"},
}

var irishBanks = []bank{
	{name: "Allied Irish Banks (AIB)", bic: "AIBKIE2DXXX", code: "AIBK931152"},
	{name: "JP Morgan Bank (Ireland)", bic: "CHASE22XXX", code: "CHAS930903"},
	{name: "Bank of Ireland", bic: "BOFIIE2DXXX", code: "BOFI900017"},
	{name: "Citibank Europe (Ireland)", bic: "CITIIE2XXXX", code: "CITI990051"},
	{name: "Bank of America Merrill Lynch International", bic: "BOFAGB22XXX", code: "BOFA990061"},
	{name: "Permanent TSB", bic: "IPBSIE2DXXX", code: "IPBS990601"},
}

var austrianBanks = []bank{
	{name: "Erste Group Bank AG", bic: "GIBAATWWXXX", code: "20111"},
	{name: "Raiffeisen Bank International AG", bic: "RZBAATWWXXX", code: "31000"},
	{name: "Bank Austria", bic: "BKAUATWWXXX", code: "12000"},
	{name: "UniCredit Bank Austria AG", bic: "BKAUATWWXXX", code: "12000"},
	{name: "Bawag P.S.K.", bic: "BWFBATW1XXX", code: "14000"},
	{name: "Volksbank", bic: "VBOEATWWXXX", code: "43000"},
	{name: "Hypo Vorarlberg Bank AG", bic: "HYPVAT2BXXX", code: "58000"},
	{name: "Oberbank AG", bic: "OBKLAT2LXXX", code: "15000"},
	{name: "Bank für Tirol und Vorarlberg AG (BTV)", bic: "BTVAAT22XXX", code: "16000"},
	{name: "Steiermärkische Bank und Sparkassen AG (StMK)", bic: "STSPAT2GXXX", code: "20815"},
}

var swissBanks = []bank{
	{name: "Credit Agricole (Suisse) SA", bic: "AGRICHZZXXX", code: "80000"},
	{name: "UBS Group AG", bic: "UBSWCHZHXXX", code: "00227"},
	{name: "Credit Suisse Group AG", bic: "CSAGCHZZXXX", code: "04835"},
	{name: "Julius Baer Group Ltd", bic: "BAERCHZZXXX", code: "08515"},
	{name: "Swiss PostFinance", bic: "POFICHBEXXX", code: "09000"},
}

var belgianBanks = []bank{
	{name: "BNP Paribas Fortis", bic: "GEBABEBB", code: "001"},
	{name: "KBC Bank", bic: "KREDBEBB", code: "432"},
	{name: "Belfius Bank", bic: "GKCCBEBB", code: "052"},
	{name: "ING Belgium", bic: "BBRUBEBB", code: "310"},
	{name: "Bpost Bank", bic: "BPOTBEB1", code: "539"},
	{name: "AXA Bank", bic: "AXABBE22", code: "753"},
	{name: "Crelan", bic: "NICA BE BB", code: "103"},
	{name: "Argenta", bic: "ARSPBE22", code: "979"},
	{name: "Beobank", bic: "CTBK BE BX", code: "539"},
	{name: "VAN LANSCHOT BELGIUM", bic: "FVLBBEBB", code: "017"},
}

var spanishBanks = []bank{
	{name: "CaixaBank", bic: "CAIXESBBXXX", code: "21000418", branchCodes: []string{"0418"}},
	{name: "Bankia", bic: "CAHMESMMXXX", code: "20386406", branchCodes: []string{"6406"}},
	{name: "Bankinter", bic: "BKBKESMMXXX", code: "01280073", branchCodes: []string{"0073"}},
	{name: "Banco Sabadell", bic: "BSABESBBXXX", code: "00810085", branchCodes: []string{"0085"}},
	{name: "Kutxabank", bic: "BASKES2BXXX", code: "20953313", branchCodes: []string{"3313"}},
}

var frenchBanks = []bank{
	{name: "BNP Paribas", bic: "BNP AFR PP", code: "30004", branchCodes: []string{"02837"}},
	{name: "Crédit Agricole", bic: "AGRIFRPP", code: "18106", branchCodes: []string{"00022"}},
	{name: "Société Générale", bic: "SOGEFRPP", code: "30003", branchCodes: []string{"03540"}},
}

var italianBanks = []bank{
	{name: "Intesa Sanpaolo", bic: "BCITITMM", code: "03069", branchCodes: []string{"15601", "15602", "15603", "15604", "15605", "15606", "15607", "15608", "15609"}},
	{name: "UniCredit", bic: "UNCRITMM", code: "02008", branchCodes: []string{"23803"}},
	{name: "Banca Monte dei Paschi di Siena", bic: "BMPSITMM", code: "01030", branchCodes: []string{"14217"}},
	{name: "Banco BPM", bic: "BAPPIT22", code: "05034", branchCodes: []string{"01636"}},
	{name: "Banca Nazionale del Lavoro", bic: "BNLIITRR", code: "01005", branchCodes: []string{"16900"}},
}
