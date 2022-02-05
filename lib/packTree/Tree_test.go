package packTree

import (
	AT "myArchHafm/lib/AuxiliaryTypes"
	"reflect"
	"testing"
)

func TestCounter_calculate1(t *testing.T) {

	tests := []struct {
		name string
		c    AT.Counter
		text string
		want AT.Counter
	}{
		// TODO: Add test cases.
		{
			name:"base_text",
			c : make(AT.Counter),
			text: "ͱbaldur’s ͱgate: ͱsiege of ͱdragonspear (с англ. — «ͱврата ͱбалдура: ͱосада ͱдраконьего ͱкопья»)" +
				" — дополнение для компьютерной ролевой игры ͱbaldur's ͱgate: ͱenhanced ͱedition, разработанное и изданное " +
				"ͱbeamdog. ͱизначально игра была выпущена в 2016 году на платформах ͱmicrosoft ͱwindows, ͱlinux и macͱoͱs, " +
				"в 2018 году стала доступна также на iͱoͱs и ͱandroid, а в 2019 году были изданы версии для игровых приставок " +
				"ͱxbox ͱone, ͱplayͱstation 4 и ͱnintendo ͱswitch.\n\nͱмасштабное дополнение являлось практически " +
				"полноценной отдельной игрой и было призвано знаменовать переход ͱbeamdog от «реставрации» классических игр " +
				"к созданию оригинального контента в рамках получившей признание серии. ͱsiege of ͱdragonspear было выпущено " +
				"спустя 17 лет после ͱbaldur’s ͱgate, заполняя пробел в событиях между ней и ͱbaldur's ͱgate ͱiͱi: ͱshadows " +
				"of ͱamn. ͱв центре сюжета дополнения — угрожающий ͱвратам ͱбалдура и другим городам ͱпобережья ͱмечей крестовый " +
				"поход, возглавляемый харизматичной молодой женщиной-паладином по имени ͱцелар ͱаргент, собирающей силы вокруг " +
				"замка ͱдраконье ͱкопьё. ͱотдельной темой сюжета стало изменение отношения окружающих к главному герою из-за набирающих " +
				"силу проявлений его божественного происхождения.\n\nͱкак и оригинальные игры, дополнение представляло" +
				" собой ролевую игру с видом от третьего лица в изометрической проекции, в которой игрок управляет командой искателей " +
				"приключений, определяя её состав. ͱот других игр серии ͱsiege of ͱdragonspear отличает более линейный сюжет, а также битвы " +
				"и другие сцены, в которых на экране одновременно присутствует несколько десятков персонажей. ͱэто стало возможным в результате " +
				"адаптации игрового движка ͱinfinity ͱengine для работы на новых компьютерах и мобильных устройствах, которая была " +
				"проделана в ходе работы над улучшенными изданиями предыдущих игр серии. ͱв игру были добавлены новые классы персонажей, " +
				"напарники и другие предметы, а также внесены доработки в интерфейс пользователя.",
			want: AT.Counter{10:4,32:260,39:2,40:1,41:1,44:16,45:2,46:11,48:3,49:4,50:3,52:1,54:1,55:1,56:1,57:1,58:4,97:23,98:7,99:4,100:16,101:22,102:6,103:13,104:3,105:19,108:6,109:5,110:19,111:21,112:4,114:12,115:18,116:11,117:5,119:4,120:3,121:2,171:2,187:2,881:59,1072:109,1073:24,1074:51,1075:35,1076:49,1077:115,1078:17,1079:20,1080:107,1081:30,1082:41,1083:64,1084:30,1085:90,1086:142,1087:43,1088:84,1089:52,1090:67,1091:31,1092:2,1093:19,1094:8,1095:10,1096:4,1097:8,1099:33,1100:19,1101:2,1102:13,1103:24,1105:2,8201:2,8212:3,8217:2},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.Calculate(tt.text); !reflect.DeepEqual(tt.c, tt.want) {
				t.Errorf("calculate() = %v, want %v", tt.c, tt.want)
			}
		})
	}
}

func TestBinaryTreeEncodeFromText(t *testing.T) {
	text :="ͱbaldur’s ͱgate: ͱsiege of ͱdragonspear (с англ. — «ͱврата ͱбалдура: ͱосада ͱдраконьего ͱкопья»)" +
	" — дополнение для компьютерной ролевой игры ͱbaldur's ͱgate: ͱenhanced ͱedition, разработанное и изданное " +
	"ͱbeamdog. ͱизначально игра была выпущена в 2016 году на платформах ͱmicrosoft ͱwindows, ͱlinux и macͱoͱs, " +
	"в 2018 году стала доступна также на iͱoͱs и ͱandroid, а в 2019 году были изданы версии для игровых приставок " +
	"ͱxbox ͱone, ͱplayͱstation 4 и ͱnintendo ͱswitch.\n\nͱмасштабное дополнение являлось практически " +
	"полноценной отдельной игрой и было призвано знаменовать переход ͱbeamdog от «реставрации» классических игр " +
	"к созданию оригинального контента в рамках получившей признание серии. ͱsiege of ͱdragonspear было выпущено " +
	"спустя 17 лет после ͱbaldur’s ͱgate, заполняя пробел в событиях между ней и ͱbaldur's ͱgate ͱiͱi: ͱshadows " +
	"of ͱamn. ͱв центре сюжета дополнения — угрожающий ͱвратам ͱбалдура и другим городам ͱпобережья ͱмечей крестовый " +
	"поход, возглавляемый харизматичной молодой женщиной-паладином по имени ͱцелар ͱаргент, собирающей силы вокруг " +
	"замка ͱдраконье ͱкопьё. ͱотдельной темой сюжета стало изменение отношения окружающих к главному герою из-за набирающих " +
	"силу проявлений его божественного происхождения.\n\nͱкак и оригинальные игры, дополнение представляло" +
	" собой ролевую игру с видом от третьего лица в изометрической проекции, в которой игрок управляет командой искателей " +
	"приключений, определяя её состав. ͱот других игр серии ͱsiege of ͱdragonspear отличает более линейный сюжет, а также битвы " +
	"и другие сцены, в которых на экране одновременно присутствует несколько десятков персонажей. ͱэто стало возможным в результате " +
	"адаптации игрового движка ͱinfinity ͱengine для работы на новых компьютерах и мобильных устройствах, которая была " +
	"проделана в ходе работы над улучшенными изданиями предыдущих игр серии. ͱв игру были добавлены новые классы персонажей, " +
	"напарники и другие предметы, а также внесены доработки в интерфейс пользователя."
	testNode := BinaryTreeEncodeFromText(text)
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		// TODO: Add test cases.
		{
			name: "base text",
			args: args{
				text: text,
			},
			want: testNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinaryTreeEncodeFromText(tt.args.text)


			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTreeEncodeFromText() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_widthTraversal(t *testing.T) {

	tests := []struct {
		name string
		args *Tree
		want string
	}{
		// TODO: Add test cases.
		{
			name: "base_test",
			args: &Tree{
				One: &Tree{
					Char: AT.Char('u'),
					},
				Zero: &Tree{
					One: &Tree{
						Char: AT.Char('i'),
					},
					Zero: &Tree{
							Char: AT.Char('w'),
							},
					},
			},
			want : "110001000111010111000100011010010010001110111",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthTraversal(tt.args); got != tt.want {
				t.Errorf("widthTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ConvertToStringByte(t *testing.T) {

	tests := []struct {
		name   string
		tree *Tree
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "base_test",
			tree: &Tree{
				One: &Tree{
					Char: AT.Char('u'),
				},
				Zero: &Tree{
					One: &Tree{
						Char: AT.Char('i'),
					},
					Zero: &Tree{
						Char: AT.Char('w'),
					},
				},
			},
			want: "0000000000101101110001000111010111000100011010010010001110111",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.tree
			if got := tr.ConvertToStringBite(); got != tt.want {
				t.Errorf("ConvertToStringByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_decodeLenHeadTree(t *testing.T) {
//	readBytes := bytes.NewReader([]byte{2,79})
//	buf := bufio.NewReader(readBytes)
//	tests := []struct {
//		name string
//		args io.Reader
//		want uint16
//	}{
//		// TODO: Add test cases.
//		{
//			name: "base Test",
//			args: buf,
//			want: uint16(591),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := decodeLenHeadTree(tt.args); got != tt.want {
//				t.Errorf("decodeLenHeadTree() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_convertBiteinChar(t *testing.T) {
	//

	bites := "01000110001100001011101000010111"

	type args struct {
		bite string
	}
	tests := []struct {
		name string
		args args
		want AT.Char
	}{
		// TODO: Add test cases.
		{
			name:"base test",
			args:args{
				bites,
			},
			want: AT.Char('t'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBiteStringInChar(tt.args.bite); got != tt.want {
				t.Errorf("convertBiteinChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertCodeBiteInLenRune(t *testing.T) {
	bites1 := []AT.Bite{AT.Bite0, AT.Bite1}
	bites2:= []AT.Bite{AT.Bite1, AT.Bite0}
	type args struct {
		bite []AT.Bite
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test len 16",
			args: args{bites1},
			want: 16,
		},
		{
			name: "test len 24",
			args: args{bites2},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCodeBiteInLenRune(tt.args.bite); got != tt.want {
				t.Errorf("convertCodeBiteInLenRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertBiteStringInChar(t *testing.T) {
	type args struct {
		bites string
	}
	tests := []struct {
		name string
		args args
		want AT.Char
	}{
		// TODO: Add test cases.
		{
			name: "base test",
			args: args{
				bites: "10010001101000101",
			},
			want: AT.Char(rune(70)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBiteStringInChar(tt.args.bites); got != tt.want {
				t.Errorf("convertBiteStringInChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeLenHeadTree(t *testing.T) {
	//readBytes := bytes.NewReader([]byte{2,79})
	//
	type args struct {
		bites []AT.Bite
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name : "Test base",
			args: args{
				bites: []AT.Bite{AT.Bite(48), AT.Bite(48), AT.Bite(48), AT.Bite(48), AT.Bite(48),
					AT.Bite(48), AT.Bite(49), AT.Bite(49), AT.Bite(49), AT.Bite(49), AT.Bite(48),
					AT.Bite(48), AT.Bite(49), AT.Bite(49), AT.Bite(48), AT.Bite(49)},
			},
			want: 973,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeLenHeadTree(tt.args.bites); got != tt.want {
				t.Errorf("decodeLenHeadTree() = %v, want %v", got, tt.want)
			}
		})
	}
}