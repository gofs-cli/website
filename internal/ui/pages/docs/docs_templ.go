// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package docs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	. "github.com/gofs-cli/website/internal/ui/components"
	"github.com/gofs-cli/website/internal/ui/pages"
)

func Docs(page templ.Component, path string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<nav class=\"absolute flex w-full items-center justify-between\"><a href=\"/\" class=\"flex items-center gap-4 p-[1rem] py-4 font-medium sm:pl-[2rem]\"><svg class=\"h-20 w-20 fill-current\" viewBox=\"533.5,272,550,455\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><path d=\"M911.45915,649.75c-2.85358,-10.51 -5.56796,-95.05 -4.72283,-147.5c0.28834,-18.01 0.26846,-32.75 -0.04971,-32.75c-2.10787,0 -28.24747,15.8 -36.18181,21.87c-2.60501,2 -5.7171,4.32 -6.9003,5.17c-2.12776,1.51 -2.1377,1.65 -0.59657,7c4.6234,16.01 6.01539,39.29 3.18169,52.89c-4.11632,19.67 -13.13443,35.02 -26.48759,45.1c-25.26463,19.07 -57.67813,15.64 -81.78939,-8.64c-31.21042,-31.45 -34.83954,-86.4 -8.02383,-121.45c8.09343,-10.58 20.35289,-19.54 32.26435,-23.58c8.29229,-2.82 27.07422,-2.53 35.79405,0.54c13.68128,4.82 25.75183,14.7 34.412,28.17l4.4842,6.96l5.93585,-4.6c8.65023,-6.7 28.00884,-19 37.19598,-23.64c4.37483,-2.21 8.05366,-4.11 8.18292,-4.22c0.11931,-0.11 1.12354,-8.94 2.21724,-19.63c4.89185,-47.7 11.16575,-82.15 19.04044,-104.58c10.86747,-30.98 25.8612,-44.34 40.25836,-35.85c8.66017,5.1 12.97534,19.99 12.10038,41.71c-0.51703,12.75 -1.30251,14.78 -5.73699,14.78c-3.73849,0 -4.88191,-1.73 -6.03527,-9.15c-1.17325,-7.52 -4.31517,-17.56 -6.46281,-20.65c-8.44143,-12.12 -22.37128,26.09 -30.18631,82.8c-2.59507,18.86 -5.26968,43.1 -4.81231,43.56c0.19886,0.21 3.25129,-0.81 6.78098,-2.25c8.98828,-3.68 15.01361,-4.16 17.728,-1.43c2.56524,2.58 2.73427,7.14 0.36788,10.02c-1.06388,1.3 -5.01117,3.06 -10.23113,4.57c-17.26068,4.99 -16.38572,4.01 -17.29051,19.53c-1.81953,31.2 -3.63906,67.98 -4.52397,91.5c-1.88913,50.09 -3.61918,82.15 -4.63334,85.75c-0.92468,3.28 -4.17597,2.05 -5.27962,-2M822.9186,592.39c9.41582,-4.44 17.07177,-13.78 21.1682,-25.79c3.03255,-8.94 3.45015,-27.35 0.85508,-38.26c-3.87769,-16.31 -11.78221,-29.42 -23.71355,-39.33c-16.4752,-13.68 -33.716,-17.8 -46.73111,-11.16c-8.45137,4.31 -15.06333,14.05 -18.8217,27.73c-2.73427,9.94 -2.49564,30.95 0.45737,41.42c4.50408,15.93 11.86175,28.21 22.45082,37.43c5.21996,4.53 8.58063,6.48 15.99795,9.24c7.48692,2.79 21.029,2.18 28.33695,-1.28M817.92732,467.52c-6.7611,-4.45 -16.12721,-6.63 -22.8883,-5.33l-5.42876,1.05l6.16453,0.13c5.73699,0.12 19.67678,3.8 23.1667,6.11c0.81531,0.54 1.93884,0.99 2.4857,0.99c0.54685,0 -1.02411,-1.32 -3.49986,-2.95M1003.84755,611.44c-16.21669,-2.4 -34.00434,-13.39 -41.89892,-25.89c-4.81231,-7.6 -2.80387,-15.38 3.95723,-15.38c2.58513,0 4.4842,1.42 11.34472,8.48c12.71683,13.11 25.60269,18.65 41.01401,17.62c19.71655,-1.31 32.12516,-16.59 28.2773,-34.82c-2.28684,-10.79 -7.61618,-19.06 -23.39539,-36.26c-18.32456,-19.97 -23.7235,-32.87 -21.27757,-50.8c3.59929,-26.3 29.371,-47.9 51.27497,-42.97c12.37877,2.79 19.78615,14.43 14.69544,23.09c-4.03677,6.89 -10.6885,5.76 -15.39144,-2.61c-2.42604,-4.34 -3.72855,-5.59 -6.77104,-6.51c-7.76532,-2.34 -16.5448,3.51 -21.52614,14.35c-2.59507,5.65 -3.86775,11.29 -3.86775,20.81c0,10.19 1.12354,6.82 4.74271,14.18c2.54535,5.18 6.52247,10.7 12.42849,17.24c15.71955,17.41 19.19953,21.94 23.3755,30.39c3.84786,7.78 4.16603,9.06 4.52397,18.2c0.31817,7.94 -0.00994,11.16 -1.70022,16.76c-7.40738,24.55 -31.52859,38.31 -59.80588,34.12M613.62286,712.55c-20.7506,-3.36 -40.53676,-15.45 -51.41417,-31.42c-6.16453,-9.04 -8.90874,-16.76 -9.48542,-26.69c-0.86502,-14.78 4.32511,-27.25 16.87292,-40.54l6.95995,-7.36l-4.71288,-5.27c-9.75388,-10.91 -17.76777,-25.17 -21.59574,-38.48c-1.96867,-6.84 -2.23713,-9.93 -2.15759,-25.29c0.0696,-14.68 0.42754,-18.79 2.24707,-25.5c3.12204,-11.53 11.85181,-28.64 19.00067,-37.23c24.24052,-29.12 62.74895,-38.5 95.92804,-23.37c18.67256,8.51 36.55964,28.03 43.25114,47.18c8.77948,25.15 5.60773,52.74 -8.31217,72.36c-7.24829,10.21 -16.56469,18.08 -27.91936,23.57c-12.22963,5.91 -18.30468,7.43 -29.82837,7.43c-9.63456,0 -16.37578,-1.35 -25.54303,-5.1l-5.80659,-2.37l-4.2058,2.13c-2.31667,1.17 -4.98134,2.81 -5.9259,3.65c-1.61073,1.42 -1.57096,1.62 0.58662,3.33c4.18591,3.31 16.06755,8.79 25.06577,11.56c8.02383,2.47 10.38027,2.74 24.77743,2.8c17.55897,0.07 21.60568,-0.74 35.01851,-6.95c5.00122,-2.32 9.01811,-5.18 13.20403,-9.41c6.57218,-6.63 9.47548,-7.55 12.92563,-4.08c2.49564,2.51 2.49564,2.53 -0.52697,16.97c-3.73849,17.77 -7.96418,31.29 -13.36311,42.72c-16.19681,34.24 -41.49126,53.77 -72.11506,55.67c-4.37483,0.28 -10.19136,0.13 -12.92563,-0.31M636.49128,693c3.82797,-0.81 10.32062,-3.12 14.41705,-5.13c9.80359,-4.8 21.4466,-16.53 27.9293,-28.13c5.89607,-10.56 10.42999,-23.79 9.46554,-27.65c-0.66617,-2.68 -0.94457,-2.81 -4.88191,-2.29c-2.29678,0.3 -7.53664,1.31 -11.63306,2.25c-8.9684,2.05 -29.38095,2.34 -38.63768,0.55c-11.11604,-2.15 -23.2363,-6.1 -33.23875,-10.84l-9.74393,-4.6l-5.49836,5.42c-6.84064,6.74 -9.78371,11.67 -11.87169,19.85c-1.37211,5.38 -1.4417,7.35 -0.44743,12.48c5.15037,26.57 34.92902,44.25 64.14094,38.09M595.84515,587.08l5.06088,-2.93l-4.89185,-6.07c-10.44987,-12.99 -15.22241,-27.56 -14.41705,-44.08c0.84514,-17.48 6.78098,-31.22 19.21941,-44.48c16.38572,-17.47 37.16615,-23.72 37.16615,-11.18c0,1.4 -2.5553,3.61 -8.69994,7.52c-15.37155,9.78 -24.13115,21.27 -27.74039,36.42c-5.47848,22.97 9.89308,49.02 32.90069,55.74c3.40043,1 9.04794,1.7 13.41282,1.66c25.4436,-0.2 44.81216,-21.56 44.81216,-49.44c0,-28.51 -17.05189,-53.55 -43.88748,-64.45c-8.2724,-3.36 -15.5903,-4.27 -25.53309,-3.17c-25.19503,2.79 -45.38884,21.06 -54.02912,48.88c-2.02833,6.55 -2.32661,9.52 -2.32661,23.5c0,17.43 1.15336,22.94 7.40738,35.41c4.71288,9.37 13.85031,20.72 15.96812,19.81c0.2784,-0.12 2.78398,-1.54 5.57791,-3.14z\"></path></svg></a><div class=\"hidden items-center justify-end gap-4 px-6 py-4 font-medium sm:flex\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Search().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<a href=\"/docs\" class=\"hover:text-gray-300 dark:hover:text-gray-600\">Docs</a> <a href=\"/blogs\" class=\"hover:text-gray-300 dark:hover:text-gray-600\">Blogs</a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Github("https://github.com/gofs-cli/website").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Theme().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div><div class=\"block sm:hidden\" x-data=\"{ show: false }\"><svg @click=\"show = !show\" class=\"m-3 h-6 w-6 fill-current hover:cursor-pointer\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\"><path d=\"M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z\"></path></svg><div x-show=\"show\" class=\"fixed bottom-0 right-0 top-0 z-[199] w-1/2 overflow-y-auto bg-gray-100 p-6 shadow-xl dark:bg-gray-900 md:top-[4rem]\" x-trap.inert.noscroll=\"show\" x-on:click.outside=\"show = false\" x-on:keydown.escape.prevent.stop=\"show = false\" x-transition:enter=\"transition ease-in-out duration-300 transform\" x-transition:enter-start=\"translate-x-full\" x-transition:enter-end=\"translate-x-0\" x-transition:leave=\"transition ease-in-out duration-300 transform\" x-transition:leave-start=\"translate-x-0\" x-transition:leave-end=\"translate-x-full\"><div class=\"flex justify-end\"><button class=\"text-gray-600 focus:outline-none\" x-on:click=\"show = ! show\"><span class=\"sr-only\">Close navigation</span> <svg class=\"h-6 w-6\" stroke=\"currentColor\" fill=\"none\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Menu(path, pages.HamburgerMenu).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</div></div></nav><div class=\"mt-[8rem] flex\"><aside class=\"hidden sm:block sm:w-56 sm:pl-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Menu(path, pages.LeftMenu).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</aside><div class=\"flex px-[1rem] sm:max-w-3xl\" listen-for-intersection-of-titles>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = page.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div><aside class=\"fixed bottom-0 right-0 hidden w-64 overflow-y-auto px-6 md:top-[8rem] xl:block\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavHighlighter().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</aside></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
