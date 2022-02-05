package EncodeTable

import (
	"fmt"
	"myArchHafm/lib/packTree"
	"reflect"
	"testing"
)

func Test_ConvertTreeToTable(t *testing.T) {
	a := packTree.BinaryTreeEncodeFromText("ͱbaldur’s ͱgate: ͱsiege of ͱdragonspear (с англ. — «ͱврата ͱбалдура: ͱосада ͱдраконьего ͱкопья»)" +
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
		"напарники и другие предметы, а также внесены доработки в интерфейс пользователя.")
	tests := []struct {
		name string
		tree *packTree.Tree
		want ETable
	}{
		// TODO: Add test cases.
		{
			name:"base text",
			tree: a,
			want: ETable{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTreeToTable(tt.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTreeToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bildTableFromTree(t *testing.T) {
	type args struct {
		cd    string
		node  *packTree.Tree
		table ETable
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "base test",
			args: args{
				cd : "",
				node: packTree.BinaryTreeEncodeFromText("aaaaabbbbcccddf"),
				table: ETable{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bildTableFromTree(tt.args.cd, tt.args.node, tt.args.table)
			fmt.Println()
		})
	}
}

func TestETable_EncodeText(t *testing.T) {
	text := "ͱbaldur’s ͱgate: ͱsiege of ͱdragonspear (с англ. — «ͱврата ͱбалдура: ͱосада ͱдраконьего ͱкопья»)" +
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


	tests := []struct {
		name string
		et   ETable
		text string
		want string
	}{
		// TODO: Add test cases.
		{
			name:"Base text",
			et: ETable{10:"011011101",32:"010",39:"0110111101",40:"01101101111",41:"01100101101",44:"0110100",45:"0110111111",46:"00010100",48:"111011000",49:"011011010",50:"110111001",52:"01100101100",54:"01101101100",55:"01101101110",56:"01101101101",57:"1101101101",58:"011011001",97:"111010",98:"10001101",99:"011001001",100:"0110101",101:"0000000",102:"11011010",103:"1000111",104:"110111000",105:"0001110",108:"11011101",109:"011001000",110:"0001111",111:"0000001",112:"011011000",114:"1101111",115:"0011000",116:"1110111",117:"11101101",119:"011001101",120:"111011001",121:"0110011001",171:"0110011000",187:"0110111110",881:"10000",1072:"1010",1073:"111000",1074:"11001",1075:"001101",1076:"11010",1077:"1001",1078:"0011001",1079:"0001011",1080:"1011",1081:"011111",1082:"000100",1083:"01110",1084:"100010",1085:"1111",1086:"0010",1087:"000001",1088:"00001",1089:"11000",1090:"00111",1091:"011110",1092:"0110010101",1093:"0001101",1094:"10001100",1095:"00010101",1096:"011011100",1097:"01100111",1099:"011000",1100:"0001100",1101:"0110111100",1102:"1101100",1103:"111001",1105:"1101101100",8201:"0110010100",8212:"110110111",8217:"0110010111"},
			text: text,
			want: " ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.EncodeTextToStringByte(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeText() = %v, want %v", got, tt.want)
			}
		})
	}
}