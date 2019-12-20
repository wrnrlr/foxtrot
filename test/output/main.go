package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/parser"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/output"
	"github.com/wrnrlr/foxtrot/util"
	"log"
)

var theme *material.Theme

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	gtx := layout.NewContext(w.Queue())
	gofont.Register()
	theme = material.NewTheme()
	theme.TextSize = unit.Sp(12)
	theme.Color.Text = util.Black
	kernel := expreduce.NewEvalState()
	expressions := []string{
		//"1/c+a^2+b^2",
		"{Graphics[{Red, Circle[0,0]}], Graphics[{Green, Rectangle[]}]}",
		"Graphics[{Rectangle[{1,1}], Red, Circle[0,0]}]",
		"Graphics[{Orange, Rectangle[{0.5,1}]}]",
		"Graphics[{Red, Circle[0,0]}]",
		"{1/2,x^2,y}",
		"Sin[x]",
		"Blue",
		"x+2",
		"1/2",
		"x^3",
		"Sqrt[2]",
		"Table[i,{i,0,100}]",
		"{}",
	}
	items := []Item{}
	for i, inTxt := range expressions {
		exOut := parser.Interp(inTxt, kernel)
		exOut = kernel.Eval(exOut)
		outTxt := formattedOutput(kernel, exOut, i)
		item := Item{exOut, " in := " + inTxt, outTxt}
		items = append(items, item)
	}
	list := &layout.List{Axis: layout.Vertical}
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			paint.ColorOp{util.Black}.Add(gtx.Ops)
			list.Layout(gtx, len(items), func(i int) {
				layout.UniformInset(unit.Sp(10)).Layout(gtx, func() {
					items[i].Layout(gtx)
				})
			})
			e.Frame(gtx.Ops)
		}
	}
}

type Item struct {
	Ex            api.Ex
	InTxt, OutTxt string
}

func (i *Item) Layout(gtx *layout.Context) {
	layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func() {
			theme.Label(unit.Sp(16), i.InTxt).Layout(gtx)
		}),
		layout.Rigid(func() {
			theme.Label(unit.Sp(16), i.OutTxt).Layout(gtx)
		}),
		layout.Rigid(func() {
			w := output.FromEx(i.Ex, gtx)
			w.Layout(gtx, theme.Shaper, fnt)
		}),
	)
}

var (
	_defaultFontSize = unit.Sp(20)
	fnt              = text.Font{Size: unit.Sp(20)}
)

func formattedOutput(es *expreduce.EvalState, res api.Ex, promptNum int) (s string) {
	isNull := false
	asSym, isSym := res.(*atoms.Symbol)
	if isSym {
		if asSym.Name == "System`Null" {
			isNull = true
		}
	}

	if !isNull {
		// Print formatted result
		specialForms := []string{
			"System`FullForm",
			"System`OutputForm",
		}
		wasSpecialForm := false
		for _, specialForm := range specialForms {
			asSpecialForm, isSpecialForm := atoms.HeadAssertion(
				res, specialForm)
			if !isSpecialForm {
				continue
			}
			if len(asSpecialForm.Parts) != 2 {
				continue
			}
			s = fmt.Sprintf(
				"Out[%d]//%s= %s\n\n",
				promptNum,
				specialForm[7:],
				asSpecialForm.Parts[1].StringForm(
					expreduce.ActualStringFormArgsFull(specialForm[7:], es)))
			wasSpecialForm = true
		}
		if !wasSpecialForm {
			s = fmt.Sprintf("Out[%d]= %s\n\n", promptNum, res.StringForm(
				expreduce.ActualStringFormArgsFull("InputForm", es)))
		}
	}
	return s
}

var Sample1 = `
Graphics[{
	Thick, Green, Rectangle[{0, -1}, {2, 1}],
	Red, Disk[],
	Blue, Circle[{2, 0}],
	Yellow, Polygon[{{2, 0}, {4, 1}, {4, -1}}],
	Purple, Arrowheads[Large], Arrow[{{4, 3/2}, {0, 3/2}, {0, 0}}],
	Black, Dashed, Line[{{-1, 0}, {4, 0}}]
}]`

var sampleData2 = `
Graphics[{
	{Directive[Opacity[1.], RGBColor[0.37, 0.5, 0.71], AbsoluteThickness[1.6]],
		Line[{{-5., 0.00418624}, {-4.978, 0.00422717}, {-4.956, 0.00426575}, {-4.934, 0.00430184}, {-4.912, 0.00433531}, {-4.89, 0.00436599}, {-4.868, 0.00439373}, {-4.846, 0.00441837}, {-4.824, 0.00443974}, {-4.802, 0.00445768}, {-4.78, 0.00447202}, {-4.758, 0.00448258}, {-4.736, 0.00448918}, {-4.714, 0.00449163}, {-4.692, 0.00448975}, {-4.67, 0.00448334}, {-4.648, 0.00447221}, {-4.626, 0.00445615}, {-4.604, 0.00443496}, {-4.582, 0.00440843}, {-4.56, 0.00437634}, {-4.538, 0.00433848}, {-4.516, 0.00429462}, {-4.494, 0.00424454}, {-4.472, 0.004188}, {-4.45, 0.00412478}, {-4.428, 0.00405463}, {-4.406, 0.00397732}, {-4.384, 0.00389259}, {-4.362, 0.0038002}, {-4.34, 0.00369989}, {-4.318, 0.00359141}, {-4.296, 0.0034745}, {-4.274, 0.00334888}, {-4.252, 0.00321431}, {-4.23, 0.0030705}, {-4.208, 0.00291718}, {-4.186, 0.00275408}, {-4.164, 0.00258092}, {-4.142, 0.00239741}, {-4.12, 0.00220328}, {-4.098, 0.00199824}, {-4.076, 0.00178199}, {-4.054, 0.00155425}, {-4.032, 0.00131474}, {-4.01, 0.00106314}, {-3.988, 0.00079918}, {-3.966, 0.000522552}, {-3.944, 0.000232965}, {-3.922, -6.98798e-05}, {-3.9, -0.000386278}, {-3.878, -0.000716526}, {-3.856, -0.00106092}, {-3.834, -0.00141976}, {-3.812, -0.00179333}, {-3.79, -0.00218193}, {-3.768, -0.00258586}, {-3.746, -0.00300539}, {-3.724, -0.00344082}, {-3.702, -0.00389243}, {-3.68, -0.00436051}, {-3.658, -0.00484533}, {-3.636, -0.00534716}, {-3.614, -0.00586627}, {-3.592, -0.00640293}, {-3.57, -0.00695739}, {-3.548, -0.00752991}, {-3.526, -0.00812074}, {-3.504, -0.0087301}, {-3.482, -0.00935823}, {-3.46, -0.0100054}, {-3.438, -0.0106717}, {-3.416, -0.0113575}, {-3.394, -0.0120628}, {-3.372, -0.012788}, {-3.35, -0.0135331}, {-3.328, -0.0142983}, {-3.306, -0.0150838}, {-3.284, -0.0158898}, {-3.262, -0.0167162}, {-3.24, -0.0175633}, {-3.218, -0.0184311}, {-3.196, -0.0193197}, {-3.174, -0.0202292}, {-3.152, -0.0211596}, {-3.13, -0.0221108}, {-3.108, -0.023083}, {-3.086, -0.0240761}, {-3.064, -0.02509}, {-3.042, -0.0261247}, {-3.02, -0.0271801}, {-2.998, -0.028256}, {-2.976, -0.0293525}, {-2.954, -0.0304691}, {-2.932, -0.0316059}, {-2.91, -0.0327625}, {-2.888, -0.0339388}, {-2.866, -0.0351343}, {-2.844, -0.0363489}, {-2.822, -0.0375821}, {-2.8, -0.0388336}, {-2.778, -0.040103}, {-2.756, -0.0413898}, {-2.734, -0.0426935}, {-2.712, -0.0440137}, {-2.69, -0.0453496}, {-2.668, -0.0467008}, {-2.646, -0.0480666}, {-2.624, -0.0494462}, {-2.602, -0.0508389}, {-2.58, -0.052244}, {-2.558, -0.0536606}, {-2.536, -0.0550878}, {-2.514, -0.0565247}, {-2.492, -0.0579703}, {-2.47, -0.0594236}, {-2.448, -0.0608835}, {-2.426, -0.0623487}, {-2.404, -0.0638182}, {-2.382, -0.0652906}, {-2.36, -0.0667647}, {-2.338, -0.068239}, {-2.316, -0.0697121}, {-2.294, -0.0711825}, {-2.272, -0.0726485}, {-2.25, -0.0741087}, {-2.228, -0.0755611}, {-2.206, -0.0770042}, {-2.184, -0.0784359}, {-2.162, -0.0798545}, {-2.14, -0.0812578}, {-2.118, -0.0826439}, {-2.096, -0.0840105}, {-2.074, -0.0853555}, {-2.052, -0.0866765}, {-2.03, -0.0879712}, {-2.008, -0.089237}, {-1.986, -0.0904715}, {-1.964, -0.091672}, {-1.942, -0.0928358}, {-1.92, -0.09396}, {-1.898, -0.0950419}, {-1.876, -0.0960784}, {-1.854, -0.0970664}, {-1.832, -0.0980028}, {-1.81, -0.0988844}, {-1.788, -0.0997078}, {-1.766, -0.10047}, {-1.744, -0.101166}, {-1.722, -0.101794}, {-1.7, -0.102349}, {-1.678, -0.102828}, {-1.656, -0.103227}, {-1.634, -0.103542}, {-1.612, -0.103768}, {-1.59, -0.103902}, {-1.568, -0.103939}, {-1.546, -0.103875}, {-1.524, -0.103705}, {-1.502, -0.103425}, {-1.48, -0.10303}, {-1.458, -0.102515}, {-1.436, -0.101876}, {-1.414, -0.101107}, {-1.392, -0.100203}, {-1.37, -0.0991599}, {-1.348, -0.0979715}, {-1.326, -0.0966327}, {-1.304, -0.0951381}, {-1.282, -0.0934821}, {-1.26, -0.091659}, {-1.238, -0.0896633}, {-1.216, -0.0874891}, {-1.194, -0.0851305}, {-1.172, -0.0825815}, {-1.15, -0.0798362}, {-1.128, -0.0768884}, {-1.106, -0.073732}, {-1.084, -0.0703605}, {-1.062, -0.0667678}, {-1.04, -0.0629474}, {-1.018, -0.0588929}, {-0.996, -0.0545977}, {-0.974, -0.0500553}, {-0.952, -0.0452591}, {-0.93, -0.0402023}, {-0.908, -0.0348782}, {-0.886, -0.0292801}, {-0.864, -0.0234013}, {-0.842, -0.0172349}, {-0.82, -0.010774}, {-0.798, -0.00401181}, {-0.776, 0.0030585}, {-0.754, 0.0104438}, {-0.732, 0.018151}, {-0.71, 0.026187}, {-0.688, 0.0345585}, {-0.666, 0.0432723}, {-0.644, 0.0523353}, {-0.622, 0.0617542}, {-0.6, 0.0715357}, {-0.578, 0.0816864}, {-0.556, 0.0922129}, {-0.534, 0.103122}, {-0.512, 0.114419}, {-0.49, 0.126111}, {-0.468, 0.138205}, {-0.446, 0.150705}, {-0.424, 0.163619}, {-0.402, 0.176952}, {-0.38, 0.19071}, {-0.358, 0.204898}, {-0.336, 0.219521}, {-0.314, 0.234584}, {-0.292, 0.250093}, {-0.27, 0.266053}, {-0.248, 0.282467}, {-0.226, 0.29934}, {-0.204, 0.316675}, {-0.182, 0.334477}, {-0.16, 0.352749}, {-0.138, 0.371493}, {-0.116, 0.390714}, {-0.094, 0.410412}, {-0.072, 0.43059}, {-0.05, 0.451249}, {-0.028, 0.472392}, {-0.006, 0.494018}, {0.016, 0.516128}, {0.038, 0.538722}, {0.06, 0.561799}, {0.082, 0.585358}, {0.104, 0.609398}, {0.126, 0.633916}, {0.148, 0.65891}, {0.17, 0.684376}, {0.192, 0.71031}, {0.214, 0.736708}, {0.236, 0.763564}, {0.258, 0.790873}, {0.28, 0.818628}, {0.302, 0.846821}, {0.324, 0.875444}, {0.346, 0.904489}, {0.368, 0.933945}, {0.39, 0.963802}, {0.412, 0.994048}, {0.434, 1.02467}, {0.456, 1.05566}, {0.478, 1.08699}, {0.5, 1.11866}, {0.522, 1.15065}, {0.544, 1.18294}, {0.566, 1.21551}, {0.588, 1.24834}, {0.61, 1.28141}, {0.632, 1.31471}, {0.654, 1.3482}, {0.676, 1.38186}, {0.698, 1.41567}, {0.72, 1.4496}, {0.742, 1.48362}, {0.764, 1.5177}, {0.786, 1.55182}, {0.808, 1.58593}, {0.83, 1.62001}, {0.852, 1.65402}, {0.874, 1.68792}, {0.896, 1.72167}, {0.918, 1.75525}, {0.94, 1.78859}, {0.962, 1.82166}, {0.984, 1.85442}, {1.006, 1.88682}, {1.028, 1.91881}, {1.05, 1.95034}, {1.072, 1.98136}, {1.094, 2.01181}, {1.116, 2.04164}, {1.138, 2.07079}, {1.16, 2.0992}, {1.182, 2.12682}, {1.204, 2.15357}, {1.226, 2.1794}, {1.248, 2.20423}, {1.27, 2.22799}, {1.292, 2.25062}, {1.314, 2.27204}, {1.336, 2.29218}, {1.358, 2.31096}, {1.38, 2.32829}, {1.402, 2.3441}, {1.424, 2.35829}, {1.446, 2.3708}, {1.468, 2.38152}, {1.49, 2.39037}, {1.512, 2.39724}, {1.534, 2.40206}, {1.556, 2.40472}, {1.578, 2.40511}, {1.6, 2.40315}, {1.622, 2.39871}, {1.644, 2.39171}, {1.666, 2.38202}, {1.688, 2.36954}, {1.71, 2.35416}, {1.732, 2.33575}, {1.754, 2.3142}, {1.776, 2.28939}, {1.798, 2.26121}, {1.82, 2.22951}, {1.842, 2.19419}, {1.864, 2.1551}, {1.886, 2.11213}, {1.908, 2.06513}, {1.93, 2.01397}, {1.952, 1.95852}, {1.974, 1.89864}, {1.996, 1.83418}, {2.018, 1.76501}, {2.04, 1.69098}, {2.062, 1.61195}, {2.084, 1.52777}, {2.106, 1.43829}, {2.128, 1.34336}, {2.15, 1.24283}, {2.172, 1.13656}, {2.194, 1.02437}, {2.216, 0.906127}, {2.238, 0.781667}, {2.26, 0.650835}, {2.282, 0.513474}, {2.304, 0.369426}, {2.326, 0.218532}, {2.348, 0.0606353}, {2.37, -0.104424}, {2.392, -0.276805}, {2.414, -0.456667}, {2.436, -0.644167}, {2.458, -0.839463}, {2.48, -1.04271}, {2.502, -1.25408}, {2.524, -1.4737}, {2.546, -1.70175}, {2.568, -1.93838}, {2.59, -2.18373}, {2.612, -2.43795}, {2.634, -2.7012}, {2.656, -2.97362}, {2.678, -3.25535}, {2.7, -3.54653}, {2.722, -3.8473}, {2.744, -4.15779}, {2.766, -4.47813}, {2.788, -4.80845}, {2.81, -5.14887}, {2.832, -5.49951}, {2.854, -5.86048}, {2.876, -6.23189}, {2.898, -6.61384}, {2.92, -7.00643}, {2.942, -7.40974}, {2.964, -7.82386}, {2.986, -8.24887}, {3.008, -8.68483}, {3.03, -9.13181}, {3.052, -9.58986}, {3.074, -10.059}, {3.096, -10.5393}, {3.118, -11.0308}, {3.14, -11.5335}, {3.162, -12.0474}, {3.184, -12.5725}, {3.206, -13.1087}, {3.228, -13.6561}, {3.25, -14.2147}, {3.272, -14.7842}, {3.294, -15.3648}, {3.316, -15.9563}, {3.338, -16.5586}, {3.36, -17.1716}, {3.382, -17.7952}, {3.404, -18.4292}, {3.426, -19.0735}, {3.448, -19.7279}, {3.47, -20.3923}, {3.492, -21.0663}, {3.514, -21.7498}, {3.536, -22.4425}, {3.558, -23.1442}, {3.58, -23.8546}, {3.602, -24.5733}, {3.624, -25.3}, {3.646, -26.0344}, {3.668, -26.776}, {3.69, -27.5245}, {3.712, -28.2795}, {3.734, -29.0404}, {3.756, -29.8068}, {3.778, -30.5782}, {3.8, -31.354}, {3.822, -32.1337}, {3.844, -32.9167}, {3.866, -33.7023}, {3.888, -34.4899}, {3.91, -35.2788}, {3.932, -36.0683}, {3.954, -36.8576}, {3.976, -37.6459}, {3.998, -38.4325}, {4.02, -39.2164}, {4.042, -39.9968}, {4.064, -40.7726}, {4.086, -41.5431}, {4.108, -42.3071}, {4.13, -43.0635}, {4.152, -43.8115}, {4.174, -44.5496}, {4.196, -45.277}, {4.218, -45.9922}, {4.24, -46.6941}, {4.262, -47.3813}, {4.284, -48.0526}, {4.306, -48.7066}, {4.328, -49.3417}, {4.35, -49.9566}, {4.372, -50.5498}, {4.394, -51.1195}, {4.416, -51.6644}, {4.438, -52.1826}, {4.46, -52.6725}, {4.482, -53.1323}, {4.504, -53.5602}, {4.526, -53.9544}, {4.548, -54.3129}, {4.57, -54.6337}, {4.592, -54.915}, {4.614, -55.1546}, {4.636, -55.3503}, {4.658, -55.5001}, {4.68, -55.6018}, {4.702, -55.6529}, {4.724, -55.6513}, {4.746, -55.5946}, {4.768, -55.4803}, {4.79, -55.3059}, {4.812, -55.069}, {4.834, -54.767}, {4.856, -54.3971}, {4.878, -53.9568}, {4.9, -53.4433}, {4.922, -52.8538}, {4.944, -52.1855}, {4.966, -51.4355}, {4.988, -50.6009}, {5.01, -49.6786}, {5.032, -48.6657}, {5.054, -47.5591}, {5.076, -46.3557}, {5.098, -45.0522}, {5.12, -43.6455}, {5.142, -42.1324}, {5.164, -40.5095}, {5.186, -38.7735}, {5.208, -36.9211}, {5.23, -34.9488}, {5.252, -32.8532}, {5.274, -30.6308}, {5.296, -28.2781}, {5.318, -25.7917}, {5.34, -23.1678}, {5.362, -20.4031}, {5.384, -17.4937}, {5.406, -14.4362}, {5.428, -11.2269}, {5.45, -7.86205}, {5.472, -4.3381}, {5.494, -0.651343}, {5.516, 3.2019}, {5.538, 7.2253}, {5.56, 11.4225}, {5.582, 15.7973}, {5.604, 20.3531}, {5.626, 25.0938}, {5.648, 30.0229}, {5.67, 35.144}, {5.692, 40.4607}, {5.714, 45.9765}, {5.736, 51.695}, {5.758, 57.6195}, {5.78, 63.7534}, {5.802, 70.1003}, {5.824, 76.6632}, {5.846, 83.4455}, {5.868, 90.4503}, {5.89, 97.6808}, {5.912, 105.14}, {5.934, 112.83}, {5.956, 120.755}, {5.978, 128.917}, {6., 137.318}}]
	}}, {
		DisplayFunction -> Identity,
		AspectRatio -> (1/(GoldenRatio)),
		Axes -> {True, True},
		AxesLabel -> {None, None},
		AxesOrigin -> {0, 0},
		DisplayFunction :> Identity,
		Frame -> {{False, False}, {False, False}},
		FrameLabel -> {{None, None}, {None, None}},
		FrameTicks -> {{Automatic, Automatic}, {Automatic, Automatic}},
		GridLines -> {None, None},
		GridLinesStyle -> Directive[GrayLevel[0.5, 0.4]],
		Method -> {"DefaultBoundaryStyle" -> Automatic,
		"DefaultMeshStyle" -> AbsolutePointSize[6],
		"ScalingFunctions" -> None},
		PlotRange -> {{-5, 6}, {-55.6529, 137.318}},
		PlotRangeClipping -> True,
		PlotRangePadding -> {{Scaled[0.02], Scaled[0.02]},{Scaled[0.05], Scaled[0.05]}},
		Ticks -> {Automatic, Automatic}
	}]
`
