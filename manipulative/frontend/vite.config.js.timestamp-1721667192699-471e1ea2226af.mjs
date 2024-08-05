// vite.config.js
import vue from "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import AutoImport from "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/node_modules/unplugin-auto-import/dist/vite.js";
import Icons from "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/node_modules/unplugin-icons/dist/vite.js";
import { NaiveUiResolver } from "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/node_modules/unplugin-vue-components/dist/resolvers.js";
import Components from "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/node_modules/unplugin-vue-components/dist/vite.js";
import { defineConfig } from "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/node_modules/vite/dist/node/index.js";
var __vite_injected_original_import_meta_url = "file:///F:/Download/GO_LANG/tiny-rdm-1.1.14/frontend/vite.config.js";
var rootPath = new URL(".", __vite_injected_original_import_meta_url).pathname;
var vite_config_default = defineConfig({
  plugins: [
    vue(),
    AutoImport({
      imports: [
        {
          "naive-ui": ["useDialog", "useMessage", "useNotification", "useLoadingBar"]
        }
      ]
    }),
    Components({
      resolvers: [NaiveUiResolver()]
    }),
    Icons()
  ],
  resolve: {
    alias: {
      "@": rootPath + "src",
      stores: rootPath + "src/stores",
      wailsjs: rootPath + "wailsjs"
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJGOlxcXFxEb3dubG9hZFxcXFxHT19MQU5HXFxcXHRpbnktcmRtLTEuMS4xNFxcXFxmcm9udGVuZFwiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiRjpcXFxcRG93bmxvYWRcXFxcR09fTEFOR1xcXFx0aW55LXJkbS0xLjEuMTRcXFxcZnJvbnRlbmRcXFxcdml0ZS5jb25maWcuanNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL0Y6L0Rvd25sb2FkL0dPX0xBTkcvdGlueS1yZG0tMS4xLjE0L2Zyb250ZW5kL3ZpdGUuY29uZmlnLmpzXCI7aW1wb3J0IHZ1ZSBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUnXG5pbXBvcnQgQXV0b0ltcG9ydCBmcm9tICd1bnBsdWdpbi1hdXRvLWltcG9ydC92aXRlJ1xuaW1wb3J0IEljb25zIGZyb20gJ3VucGx1Z2luLWljb25zL3ZpdGUnXG5pbXBvcnQgeyBOYWl2ZVVpUmVzb2x2ZXIgfSBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy9yZXNvbHZlcnMnXG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlJ1xuaW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcblxuY29uc3Qgcm9vdFBhdGggPSBuZXcgVVJMKCcuJywgaW1wb3J0Lm1ldGEudXJsKS5wYXRobmFtZVxuLy8gaHR0cHM6Ly92aXRlanMuZGV2L2NvbmZpZy9cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XG4gICAgcGx1Z2luczogW1xuICAgICAgICB2dWUoKSxcbiAgICAgICAgQXV0b0ltcG9ydCh7XG4gICAgICAgICAgICBpbXBvcnRzOiBbXG4gICAgICAgICAgICAgICAge1xuICAgICAgICAgICAgICAgICAgICAnbmFpdmUtdWknOiBbJ3VzZURpYWxvZycsICd1c2VNZXNzYWdlJywgJ3VzZU5vdGlmaWNhdGlvbicsICd1c2VMb2FkaW5nQmFyJ10sXG4gICAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIF0sXG4gICAgICAgIH0pLFxuICAgICAgICBDb21wb25lbnRzKHtcbiAgICAgICAgICAgIHJlc29sdmVyczogW05haXZlVWlSZXNvbHZlcigpXSxcbiAgICAgICAgfSksXG4gICAgICAgIEljb25zKCksXG4gICAgXSxcbiAgICByZXNvbHZlOiB7XG4gICAgICAgIGFsaWFzOiB7XG4gICAgICAgICAgICAnQCc6IHJvb3RQYXRoICsgJ3NyYycsXG4gICAgICAgICAgICBzdG9yZXM6IHJvb3RQYXRoICsgJ3NyYy9zdG9yZXMnLFxuICAgICAgICAgICAgd2FpbHNqczogcm9vdFBhdGggKyAnd2FpbHNqcycsXG4gICAgICAgIH0sXG4gICAgfSxcbn0pXG4iXSwKICAibWFwcGluZ3MiOiAiO0FBQWdVLE9BQU8sU0FBUztBQUNoVixPQUFPLGdCQUFnQjtBQUN2QixPQUFPLFdBQVc7QUFDbEIsU0FBUyx1QkFBdUI7QUFDaEMsT0FBTyxnQkFBZ0I7QUFDdkIsU0FBUyxvQkFBb0I7QUFMNEssSUFBTSwyQ0FBMkM7QUFPMVAsSUFBTSxXQUFXLElBQUksSUFBSSxLQUFLLHdDQUFlLEVBQUU7QUFFL0MsSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDeEIsU0FBUztBQUFBLElBQ0wsSUFBSTtBQUFBLElBQ0osV0FBVztBQUFBLE1BQ1AsU0FBUztBQUFBLFFBQ0w7QUFBQSxVQUNJLFlBQVksQ0FBQyxhQUFhLGNBQWMsbUJBQW1CLGVBQWU7QUFBQSxRQUM5RTtBQUFBLE1BQ0o7QUFBQSxJQUNKLENBQUM7QUFBQSxJQUNELFdBQVc7QUFBQSxNQUNQLFdBQVcsQ0FBQyxnQkFBZ0IsQ0FBQztBQUFBLElBQ2pDLENBQUM7QUFBQSxJQUNELE1BQU07QUFBQSxFQUNWO0FBQUEsRUFDQSxTQUFTO0FBQUEsSUFDTCxPQUFPO0FBQUEsTUFDSCxLQUFLLFdBQVc7QUFBQSxNQUNoQixRQUFRLFdBQVc7QUFBQSxNQUNuQixTQUFTLFdBQVc7QUFBQSxJQUN4QjtBQUFBLEVBQ0o7QUFDSixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=
