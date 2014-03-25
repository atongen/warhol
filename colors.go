package main

type Colors map[string]string
type Palette map[string]Colors

var (
	colors Colors
  palettes Palette = Palette{
    "high": Colors{
      "000": "FF0000,BF3030,A60000,FF4040,FF7373,009999,1D7373,006363,33CCCC,5CCCCC,9FEE00,86B32D,679B00,B9F73E,C9F76F",
      "030": "FF7400,BF7130,A64B00,FF9640,FFB273,1240AB,2A4480,06266F,4671D5,6C8CD5,00CC00,269926,008500,39E639,67E667",
      "060": "FFAA00,BF8F30,A66F00,FFBF40,FFD073,3914AF,412C84,200772,6A48D7,876ED7,009999,1D7373,006363,33CCCC,5CCCCC",
      "090": "FFD300,BFA730,A68900,FFDE40,FFE773,7109AA,5F2580,48036F,9F3ED5,AD66D5,1240AB,2A4480,06266F,4671D5,6C8CD5",
      "120": "FFFF00,BFBF30,A6A600,FFFF40,FFFF73,CD0074,992667,85004B,E6399B,E667AF,3914AF,412C84,200772,6A48D7,876ED7",
      "150": "9FEE00,86B32D,679B00,B9F73E,C9F76F,FF0000,BF3030,A60000,FF4040,FF7373,7109AA,5F2580,48036F,9F3ED5,AD66D5",
      "180": "00CC00,269926,008500,39E639,67E667,FF7400,BF7130,A64B00,FF9640,FFB273,CD0074,992667,85004B,E6399B,E667AF",
      "210": "009999,1D7373,006363,33CCCC,5CCCCC,FFAA00,BF8F30,A66F00,FFBF40,FFD073,FF0000,BF3030,A60000,FF4040,FF7373",
      "240": "1240AB,2A4480,06266F,4671D5,6C8CD5,FFD300,BFA730,A68900,FFDE40,FFE773,FF7400,BF7130,A64B00,FF9640,FFB273",
      "270": "3914AF,412C84,200772,6A48D7,876ED7,FFFF00,BFBF30,A6A600,FFFF40,FFFF73,FFAA00,BF8F30,A66F00,FFBF40,FFD073",
      "300": "7109AA,5F2580,48036F,9F3ED5,AD66D5,9FEE00,86B32D,679B00,B9F73E,C9F76F,FFD300,BFA730,A68900,FFDE40,FFE773",
      "330": "CD0074,992667,85004B,E6399B,E667AF,00CC00,269926,008500,39E639,67E667,FFFF00,BFBF30,A6A600,FFFF40,FFFF73",
    },
    "low": Colors{
      "000": "FF0000,009999,9FEE00",
      "030": "FF7400,1240AB,00CC00",
      "060": "FFAA00,3914AF,009999",
      "090": "FFD300,7109AA,1240AB",
      "120": "FFFF00,CD0074,3914AF",
      "150": "9FEE00,FF0000,7109AA",
      "180": "00CC00,FF7400,CD0074",
      "210": "009999,FFAA00,FF0000",
      "240": "1240AB,FFD300,FF7400",
      "270": "3914AF,FFFF00,FFAA00",
      "300": "7109AA,9FEE00,FFD300",
      "330": "CD0074,00CC00,FFFF00",
    },
  }
)
