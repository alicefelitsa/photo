/** EasyWeb iframe v3.1.8 date:2020-05-04 License By http://easyweb.vip */
layui.config({  // common.js是配置layui扩展模块的目录，每个页面都需要引入
    pageTabs: false,                // 关闭多标签
    defaultTheme: 'theme-admin', // 默认主题
    cacheTab: true,                // 刷新页面不恢复已经打开的Tab
    closeFooter: true,              // 是否关闭页脚
    tabAutoRefresh: true,           // 是否切换Tab自动刷新页面
    getAjaxHeaders: function(){},   // ajax统一传递header
    ajaxSuccessBefore: function(){},// ajax统一预处理
    version: true,                  // 加载模块不缓存，更新组件缓存，设为true不缓存，也可以设一个固定值
    base: getProjectUrl() + 'assets/module/'
}).extend({
    steps: 'steps/steps',
    notice: 'notice/notice',
    cascader: 'cascader/cascader',
    dropdown: 'dropdown/dropdown',
    fileChoose: 'fileChoose/fileChoose',
    Split: 'Split/Split',
    Cropper: 'Cropper/Cropper',
    tagsInput: 'tagsInput/tagsInput',
    citypicker: 'city-picker/city-picker',
    introJs: 'introJs/introJs',
    zTree: 'zTree/zTree'
}).use(['layer', 'admin'], function () {
    var $ = layui.jquery;
    var layer = layui.layer;
    var admin = layui.admin;

});

/** 获取当前项目的根路径，通过获取layui.js全路径截取assets之前的地址 */
function getProjectUrl() {
    var layuiDir = layui.cache.dir;
    if (!layuiDir) {
        var js = document.scripts, last = js.length - 1, src;
        for (var i = last; i > 0; i--) {
            if (js[i].readyState === 'interactive') {
                src = js[i].src;
                break;
            }
        }
        var jsPath = src || js[last].src;
        layuiDir = jsPath.substring(0, jsPath.lastIndexOf('/') + 1);
    }
    return layuiDir.substring(0, layuiDir.indexOf('assets'));
}