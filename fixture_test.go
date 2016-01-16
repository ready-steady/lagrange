package lagrange

var trainPoints []float64 = []float64{
	0.0000000000000000e+00, 0.0000000000000000e+00, 0.0000000000000000e+00,
	1.4644660940672627e-01, 0.0000000000000000e+00, 0.0000000000000000e+00,
	5.0000000000000000e-01, 0.0000000000000000e+00, 0.0000000000000000e+00,
	8.5355339059327373e-01, 0.0000000000000000e+00, 0.0000000000000000e+00,
	1.0000000000000000e+00, 0.0000000000000000e+00, 0.0000000000000000e+00,
	0.0000000000000000e+00, 1.4644660940672627e-01, 0.0000000000000000e+00,
	1.4644660940672627e-01, 1.4644660940672627e-01, 0.0000000000000000e+00,
	5.0000000000000000e-01, 1.4644660940672627e-01, 0.0000000000000000e+00,
	8.5355339059327373e-01, 1.4644660940672627e-01, 0.0000000000000000e+00,
	1.0000000000000000e+00, 1.4644660940672627e-01, 0.0000000000000000e+00,
	0.0000000000000000e+00, 5.0000000000000000e-01, 0.0000000000000000e+00,
	1.4644660940672627e-01, 5.0000000000000000e-01, 0.0000000000000000e+00,
	5.0000000000000000e-01, 5.0000000000000000e-01, 0.0000000000000000e+00,
	8.5355339059327373e-01, 5.0000000000000000e-01, 0.0000000000000000e+00,
	1.0000000000000000e+00, 5.0000000000000000e-01, 0.0000000000000000e+00,
	0.0000000000000000e+00, 8.5355339059327373e-01, 0.0000000000000000e+00,
	1.4644660940672627e-01, 8.5355339059327373e-01, 0.0000000000000000e+00,
	5.0000000000000000e-01, 8.5355339059327373e-01, 0.0000000000000000e+00,
	8.5355339059327373e-01, 8.5355339059327373e-01, 0.0000000000000000e+00,
	1.0000000000000000e+00, 8.5355339059327373e-01, 0.0000000000000000e+00,
	0.0000000000000000e+00, 1.0000000000000000e+00, 0.0000000000000000e+00,
	1.4644660940672627e-01, 1.0000000000000000e+00, 0.0000000000000000e+00,
	5.0000000000000000e-01, 1.0000000000000000e+00, 0.0000000000000000e+00,
	8.5355339059327373e-01, 1.0000000000000000e+00, 0.0000000000000000e+00,
	1.0000000000000000e+00, 1.0000000000000000e+00, 0.0000000000000000e+00,
	0.0000000000000000e+00, 0.0000000000000000e+00, 1.4644660940672627e-01,
	1.4644660940672627e-01, 0.0000000000000000e+00, 1.4644660940672627e-01,
	5.0000000000000000e-01, 0.0000000000000000e+00, 1.4644660940672627e-01,
	8.5355339059327373e-01, 0.0000000000000000e+00, 1.4644660940672627e-01,
	1.0000000000000000e+00, 0.0000000000000000e+00, 1.4644660940672627e-01,
	0.0000000000000000e+00, 1.4644660940672627e-01, 1.4644660940672627e-01,
	1.4644660940672627e-01, 1.4644660940672627e-01, 1.4644660940672627e-01,
	5.0000000000000000e-01, 1.4644660940672627e-01, 1.4644660940672627e-01,
	8.5355339059327373e-01, 1.4644660940672627e-01, 1.4644660940672627e-01,
	1.0000000000000000e+00, 1.4644660940672627e-01, 1.4644660940672627e-01,
	0.0000000000000000e+00, 5.0000000000000000e-01, 1.4644660940672627e-01,
	1.4644660940672627e-01, 5.0000000000000000e-01, 1.4644660940672627e-01,
	5.0000000000000000e-01, 5.0000000000000000e-01, 1.4644660940672627e-01,
	8.5355339059327373e-01, 5.0000000000000000e-01, 1.4644660940672627e-01,
	1.0000000000000000e+00, 5.0000000000000000e-01, 1.4644660940672627e-01,
	0.0000000000000000e+00, 8.5355339059327373e-01, 1.4644660940672627e-01,
	1.4644660940672627e-01, 8.5355339059327373e-01, 1.4644660940672627e-01,
	5.0000000000000000e-01, 8.5355339059327373e-01, 1.4644660940672627e-01,
	8.5355339059327373e-01, 8.5355339059327373e-01, 1.4644660940672627e-01,
	1.0000000000000000e+00, 8.5355339059327373e-01, 1.4644660940672627e-01,
	0.0000000000000000e+00, 1.0000000000000000e+00, 1.4644660940672627e-01,
	1.4644660940672627e-01, 1.0000000000000000e+00, 1.4644660940672627e-01,
	5.0000000000000000e-01, 1.0000000000000000e+00, 1.4644660940672627e-01,
	8.5355339059327373e-01, 1.0000000000000000e+00, 1.4644660940672627e-01,
	1.0000000000000000e+00, 1.0000000000000000e+00, 1.4644660940672627e-01,
	0.0000000000000000e+00, 0.0000000000000000e+00, 5.0000000000000000e-01,
	1.4644660940672627e-01, 0.0000000000000000e+00, 5.0000000000000000e-01,
	5.0000000000000000e-01, 0.0000000000000000e+00, 5.0000000000000000e-01,
	8.5355339059327373e-01, 0.0000000000000000e+00, 5.0000000000000000e-01,
	1.0000000000000000e+00, 0.0000000000000000e+00, 5.0000000000000000e-01,
	0.0000000000000000e+00, 1.4644660940672627e-01, 5.0000000000000000e-01,
	1.4644660940672627e-01, 1.4644660940672627e-01, 5.0000000000000000e-01,
	5.0000000000000000e-01, 1.4644660940672627e-01, 5.0000000000000000e-01,
	8.5355339059327373e-01, 1.4644660940672627e-01, 5.0000000000000000e-01,
	1.0000000000000000e+00, 1.4644660940672627e-01, 5.0000000000000000e-01,
	0.0000000000000000e+00, 5.0000000000000000e-01, 5.0000000000000000e-01,
	1.4644660940672627e-01, 5.0000000000000000e-01, 5.0000000000000000e-01,
	5.0000000000000000e-01, 5.0000000000000000e-01, 5.0000000000000000e-01,
	8.5355339059327373e-01, 5.0000000000000000e-01, 5.0000000000000000e-01,
	1.0000000000000000e+00, 5.0000000000000000e-01, 5.0000000000000000e-01,
	0.0000000000000000e+00, 8.5355339059327373e-01, 5.0000000000000000e-01,
	1.4644660940672627e-01, 8.5355339059327373e-01, 5.0000000000000000e-01,
	5.0000000000000000e-01, 8.5355339059327373e-01, 5.0000000000000000e-01,
	8.5355339059327373e-01, 8.5355339059327373e-01, 5.0000000000000000e-01,
	1.0000000000000000e+00, 8.5355339059327373e-01, 5.0000000000000000e-01,
	0.0000000000000000e+00, 1.0000000000000000e+00, 5.0000000000000000e-01,
	1.4644660940672627e-01, 1.0000000000000000e+00, 5.0000000000000000e-01,
	5.0000000000000000e-01, 1.0000000000000000e+00, 5.0000000000000000e-01,
	8.5355339059327373e-01, 1.0000000000000000e+00, 5.0000000000000000e-01,
	1.0000000000000000e+00, 1.0000000000000000e+00, 5.0000000000000000e-01,
	0.0000000000000000e+00, 0.0000000000000000e+00, 8.5355339059327373e-01,
	1.4644660940672627e-01, 0.0000000000000000e+00, 8.5355339059327373e-01,
	5.0000000000000000e-01, 0.0000000000000000e+00, 8.5355339059327373e-01,
	8.5355339059327373e-01, 0.0000000000000000e+00, 8.5355339059327373e-01,
	1.0000000000000000e+00, 0.0000000000000000e+00, 8.5355339059327373e-01,
	0.0000000000000000e+00, 1.4644660940672627e-01, 8.5355339059327373e-01,
	1.4644660940672627e-01, 1.4644660940672627e-01, 8.5355339059327373e-01,
	5.0000000000000000e-01, 1.4644660940672627e-01, 8.5355339059327373e-01,
	8.5355339059327373e-01, 1.4644660940672627e-01, 8.5355339059327373e-01,
	1.0000000000000000e+00, 1.4644660940672627e-01, 8.5355339059327373e-01,
	0.0000000000000000e+00, 5.0000000000000000e-01, 8.5355339059327373e-01,
	1.4644660940672627e-01, 5.0000000000000000e-01, 8.5355339059327373e-01,
	5.0000000000000000e-01, 5.0000000000000000e-01, 8.5355339059327373e-01,
	8.5355339059327373e-01, 5.0000000000000000e-01, 8.5355339059327373e-01,
	1.0000000000000000e+00, 5.0000000000000000e-01, 8.5355339059327373e-01,
	0.0000000000000000e+00, 8.5355339059327373e-01, 8.5355339059327373e-01,
	1.4644660940672627e-01, 8.5355339059327373e-01, 8.5355339059327373e-01,
	5.0000000000000000e-01, 8.5355339059327373e-01, 8.5355339059327373e-01,
	8.5355339059327373e-01, 8.5355339059327373e-01, 8.5355339059327373e-01,
	1.0000000000000000e+00, 8.5355339059327373e-01, 8.5355339059327373e-01,
	0.0000000000000000e+00, 1.0000000000000000e+00, 8.5355339059327373e-01,
	1.4644660940672627e-01, 1.0000000000000000e+00, 8.5355339059327373e-01,
	5.0000000000000000e-01, 1.0000000000000000e+00, 8.5355339059327373e-01,
	8.5355339059327373e-01, 1.0000000000000000e+00, 8.5355339059327373e-01,
	1.0000000000000000e+00, 1.0000000000000000e+00, 8.5355339059327373e-01,
	0.0000000000000000e+00, 0.0000000000000000e+00, 1.0000000000000000e+00,
	1.4644660940672627e-01, 0.0000000000000000e+00, 1.0000000000000000e+00,
	5.0000000000000000e-01, 0.0000000000000000e+00, 1.0000000000000000e+00,
	8.5355339059327373e-01, 0.0000000000000000e+00, 1.0000000000000000e+00,
	1.0000000000000000e+00, 0.0000000000000000e+00, 1.0000000000000000e+00,
	0.0000000000000000e+00, 1.4644660940672627e-01, 1.0000000000000000e+00,
	1.4644660940672627e-01, 1.4644660940672627e-01, 1.0000000000000000e+00,
	5.0000000000000000e-01, 1.4644660940672627e-01, 1.0000000000000000e+00,
	8.5355339059327373e-01, 1.4644660940672627e-01, 1.0000000000000000e+00,
	1.0000000000000000e+00, 1.4644660940672627e-01, 1.0000000000000000e+00,
	0.0000000000000000e+00, 5.0000000000000000e-01, 1.0000000000000000e+00,
	1.4644660940672627e-01, 5.0000000000000000e-01, 1.0000000000000000e+00,
	5.0000000000000000e-01, 5.0000000000000000e-01, 1.0000000000000000e+00,
	8.5355339059327373e-01, 5.0000000000000000e-01, 1.0000000000000000e+00,
	1.0000000000000000e+00, 5.0000000000000000e-01, 1.0000000000000000e+00,
	0.0000000000000000e+00, 8.5355339059327373e-01, 1.0000000000000000e+00,
	1.4644660940672627e-01, 8.5355339059327373e-01, 1.0000000000000000e+00,
	5.0000000000000000e-01, 8.5355339059327373e-01, 1.0000000000000000e+00,
	8.5355339059327373e-01, 8.5355339059327373e-01, 1.0000000000000000e+00,
	1.0000000000000000e+00, 8.5355339059327373e-01, 1.0000000000000000e+00,
	0.0000000000000000e+00, 1.0000000000000000e+00, 1.0000000000000000e+00,
	1.4644660940672627e-01, 1.0000000000000000e+00, 1.0000000000000000e+00,
	5.0000000000000000e-01, 1.0000000000000000e+00, 1.0000000000000000e+00,
	8.5355339059327373e-01, 1.0000000000000000e+00, 1.0000000000000000e+00,
	1.0000000000000000e+00, 1.0000000000000000e+00, 1.0000000000000000e+00,
}

var trainValues []float64 = []float64{
	0.0000000000000000e+00,
	1.4592370657462522e-01,
	4.7942553860420301e-01,
	7.5362083505949129e-01,
	8.4147098480789650e-01,
	1.4592370657462522e-01,
	2.0562937387258196e-01,
	4.9775237746795942e-01,
	7.6175998141628920e-01,
	8.4718609835094671e-01,
	4.7942553860420301e-01,
	4.9775237746795942e-01,
	6.4963693908006248e-01,
	8.3559696382834925e-01,
	8.9924214660378510e-01,
	7.5362083505949129e-01,
	7.6175998141628920e-01,
	8.3559696382834925e-01,
	9.3459072472042048e-01,
	9.6739746993765774e-01,
	8.4147098480789650e-01,
	8.4718609835094671e-01,
	8.9924214660378510e-01,
	9.6739746993765774e-01,
	9.8776594599273559e-01,
	1.4592370657462522e-01,
	2.0562937387258196e-01,
	4.9775237746795942e-01,
	7.6175998141628920e-01,
	8.4718609835094671e-01,
	2.0562937387258196e-01,
	2.5094170684085071e-01,
	5.1516152911195368e-01,
	7.6966759115938710e-01,
	8.5274664352439244e-01,
	4.9775237746795942e-01,
	5.1516152911195368e-01,
	6.6097146534071427e-01,
	8.4147098480789650e-01,
	9.0337891803073367e-01,
	7.6175998141628920e-01,
	7.6966759115938710e-01,
	8.4147098480789650e-01,
	9.3770260587034593e-01,
	9.6942476090417373e-01,
	8.4718609835094671e-01,
	8.5274664352439244e-01,
	9.0337891803073367e-01,
	9.6942476090417373e-01,
	9.8891698386577165e-01,
	4.7942553860420301e-01,
	4.9775237746795942e-01,
	6.4963693908006248e-01,
	8.3559696382834925e-01,
	8.9924214660378510e-01,
	4.9775237746795942e-01,
	5.1516152911195368e-01,
	6.6097146534071427e-01,
	8.4147098480789650e-01,
	9.0337891803073367e-01,
	6.4963693908006248e-01,
	6.6097146534071427e-01,
	7.6175998141628920e-01,
	8.9498665067142025e-01,
	9.4071933374144412e-01,
	8.3559696382834925e-01,
	8.4147098480789650e-01,
	8.9498665067142025e-01,
	9.6529300687171882e-01,
	9.8655177519934611e-01,
	8.9924214660378510e-01,
	9.0337891803073367e-01,
	9.4071933374144412e-01,
	9.8655177519934611e-01,
	9.9749498660405445e-01,
	7.5362083505949129e-01,
	7.6175998141628920e-01,
	8.3559696382834925e-01,
	9.3459072472042048e-01,
	9.6739746993765774e-01,
	7.6175998141628920e-01,
	7.6966759115938710e-01,
	8.4147098480789650e-01,
	9.3770260587034582e-01,
	9.6942476090417373e-01,
	8.3559696382834925e-01,
	8.4147098480789650e-01,
	8.9498665067142025e-01,
	9.6529300687171882e-01,
	9.8655177519934611e-01,
	9.3459072472042048e-01,
	9.3770260587034582e-01,
	9.6529300687171882e-01,
	9.9573429593068552e-01,
	9.9999462012099771e-01,
	9.6739746993765774e-01,
	9.6942476090417373e-01,
	9.8655177519934611e-01,
	9.9999462012099771e-01,
	9.9671829756338270e-01,
	8.4147098480789650e-01,
	8.4718609835094671e-01,
	8.9924214660378510e-01,
	9.6739746993765774e-01,
	9.8776594599273559e-01,
	8.4718609835094671e-01,
	8.5274664352439244e-01,
	9.0337891803073367e-01,
	9.6942476090417373e-01,
	9.8891698386577165e-01,
	8.9924214660378510e-01,
	9.0337891803073367e-01,
	9.4071933374144412e-01,
	9.8655177519934611e-01,
	9.9749498660405445e-01,
	9.6739746993765774e-01,
	9.6942476090417373e-01,
	9.8655177519934611e-01,
	9.9999462012099771e-01,
	9.9671829756338270e-01,
	9.8776594599273559e-01,
	9.8891698386577165e-01,
	9.9749498660405445e-01,
	9.9671829756338270e-01,
	9.8702664499035375e-01,
}

var testPoints []float64 = []float64{
	2.1841829698237580e-01, 9.5631757650892968e-01, 8.2950923393270060e-01,
	5.6169544276689420e-01, 4.1530708147596124e-01, 6.6118734916031333e-02,
	2.5757779238204492e-01, 1.0995679353247061e-01, 4.3828997779107100e-02,
	6.3396571230241250e-01, 6.1727229065601817e-02, 4.4953896030717627e-01,
	4.0130628149761666e-01, 7.5467348642141019e-01, 7.9728695410625616e-01,
	1.8383711584052001e-03, 8.9750406089973112e-01, 3.5075233798054661e-01,
	9.4544750211680198e-02, 1.3616891583461951e-02, 8.5909685527985680e-01,
	8.4084745065566813e-01, 1.2310391576455423e-01, 7.5123640739642129e-03,
	2.6030299776736493e-01, 9.1248370698143277e-01, 1.1366404644249929e-01,
	3.5162865990389613e-01, 8.2288731668591963e-01, 2.6713227026017933e-01,
	6.9206649978347967e-01, 5.6166247487906173e-01, 8.6121579062360976e-01,
	4.5379377501195045e-01, 9.1197702838508465e-01, 5.9791687714408193e-01,
	1.8895469101496784e-01, 7.6149205615519266e-01, 3.9698847585915975e-01,
	1.8531411707601886e-01, 5.7436586101970610e-01, 3.6702666772782006e-01,
	6.1720482704566904e-01, 3.6152870409216176e-01, 2.1292999768107831e-01,
	7.1447121474588438e-01, 1.1770686791481653e-01, 2.9932914872905197e-01,
	8.2500295468209583e-01, 8.2466007384060969e-01, 6.1861770718754039e-02,
	7.1078052494219079e-01, 8.8283333963843108e-02, 7.7799400859037671e-01,
	7.4530306861774076e-01, 3.0867491954643533e-01, 8.9937309073164395e-01,
	7.6353672457718613e-01, 7.6173064609774654e-01, 4.0696964057201429e-01,
}

var testValues []float64 = []float64{
	9.5947556494737563e-01,
	6.4586728715819985e-01,
	2.7534879626348424e-01,
	7.0327543929107572e-01,
	9.2013360586250625e-01,
	8.2100257262211063e-01,
	7.6036669384598421e-01,
	7.5138236491533672e-01,
	8.1620741104092853e-01,
	8.0459734472798372e-01,
	9.4532249100399723e-01,
	9.2506575677864133e-01,
	7.7095310539338369e-01,
	6.4998203550886924e-01,
	6.7999720429763522e-01,
	7.0754728243242926e-01,
	9.1979843986291809e-01,
	8.7040695343527485e-01,
	9.3502427317572523e-01,
	9.1365245295168906e-01,
}
