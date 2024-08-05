<script setup>
import ContentPane from './components/content/ContentPane.vue'
import BrowserPane from './components/sidebar/BrowserPane.vue'
import { computed, onMounted, reactive, ref, watchEffect } from 'vue'
import { debounce } from 'lodash'
import { useThemeVars } from 'naive-ui'
import Ribbon from './components/sidebar/Ribbon.vue'
import ConnectionPane from './components/sidebar/ConnectionPane.vue'
import ContentServerPane from './components/content/ContentServerPane.vue'
import useTabStore from './stores/tab.js'
import usePreferencesStore from './stores/preferences.js'
import ContentLogPane from './components/content/ContentLogPane.vue'
import ContentValueTab from '@/components/content/ContentValueTab.vue'
import ToolbarControlWidget from '@/components/common/ToolbarControlWidget.vue'
import { EventsOn, WindowIsFullscreen, WindowIsMaximised, WindowToggleMaximise } from 'wailsjs/runtime/runtime.js'
import { isMacOS } from '@/utils/platform.js'
import iconUrl from '@/assets/images/icon.png'
import ResizeableWrapper from '@/components/common/ResizeableWrapper.vue'
import { extraTheme } from '@/utils/extra_theme.js'

const themeVars = useThemeVars()

const props = defineProps({
    loading: Boolean,
})

const data = reactive({
    navMenuWidth: 50,
    toolbarHeight: 38,
})

const tabStore = useTabStore()
const prefStore = usePreferencesStore()
const logPaneRef = ref(null)
const exThemeVars = computed(() => {
    return extraTheme(prefStore.isDark)
})
// const preferences = ref({})
// provide('preferences', preferences)

const saveSidebarWidth = debounce(prefStore.savePreferences, 1000, { trailing: true })
const handleResize = () => {
    saveSidebarWidth()
}

watchEffect(() => {
    if (tabStore.nav === 'log') {
        logPaneRef.value?.refresh()
    }
})

const logoWrapperWidth = computed(() => {
    return `${data.navMenuWidth + prefStore.behavior.asideWidth - 4}px`
})

const logoPaddingLeft = ref(10)
const maximised = ref(false)
const hideRadius = ref(false)
const wrapperStyle = computed(() => {
    return hideRadius.value
        ? {}
        : {
            border: `1px solid ${themeVars.value.borderColor}`,
            borderRadius: '10px',
        }
})
const spinStyle = computed(() => {
    return hideRadius.value
        ? {
            backgroundColor: themeVars.value.bodyColor,
        }
        : {
            backgroundColor: themeVars.value.bodyColor,
            borderRadius: '10px',
        }
})

const onToggleFullscreen = (fullscreen) => {
    hideRadius.value = fullscreen
    if (fullscreen) {
        logoPaddingLeft.value = 10
    } else {
        logoPaddingLeft.value = isMacOS() ? 70 : 10
    }
}

const onToggleMaximize = (isMaximised) => {
    if (isMaximised) {
        maximised.value = true
        if (!isMacOS()) {
            hideRadius.value = true
        }
    } else {
        maximised.value = false
        if (!isMacOS()) {
            hideRadius.value = false
        }
    }
}

EventsOn('window_changed', (info) => {
    const { fullscreen, maximised } = info
    onToggleFullscreen(fullscreen === true)
    onToggleMaximize(maximised)
})

EventsOn('login_error', (info) => {
    console.log("ERROR")
})

onMounted(async () => {
    const fullscreen = await WindowIsFullscreen()
    onToggleFullscreen(fullscreen === true)
    const maximised = await WindowIsMaximised()
    onToggleMaximize(maximised)
})

const onKeyShortcut = (e) => {
    switch (e.key) {
        case 'w':
            if (e.metaKey) {
                // close current tab
                const tabStore = useTabStore()
                const currentTab = tabStore.currentTab
                if (currentTab != null) {
                    tabStore.closeTab(currentTab.name)
                }
            }
            break
    }
}
</script>

<template>
    <!-- app content-->
    <n-spin :show="props.loading" :style="spinStyle" :theme-overrides="{ opacitySpinning: 0 }">
        <div id="app-content-wrapper" :style="wrapperStyle" class="flex-box-v" tabindex="0" @keydown="onKeyShortcut">
            <router-view />
        </div>
    </n-spin>
</template>

<style lang="scss" scoped>
#app-content-wrapper {
    width: 100vw;
    height: 100vh;
    overflow: hidden;
    box-sizing: border-box;
    background-color: v-bind('themeVars.bodyColor');
    color: v-bind('themeVars.textColorBase');

    #app-toolbar {
        background-color: v-bind('exThemeVars.titleColor');
        border-bottom: 1px solid v-bind('exThemeVars.splitColor');

        &-title {
            padding-left: 10px;
            padding-right: 10px;
            box-sizing: border-box;
            align-self: center;
            align-items: baseline;
        }
    }

    .app-toolbar-tab {
        align-self: flex-end;
        margin-bottom: -1px;
        margin-left: 3px;
        overflow: auto;
    }

    #app-content {
        height: calc(100% - 60px);

        .content-area {
            overflow: hidden;
        }
    }

    .app-side {
        //overflow: hidden;
        height: 100%;
        background-color: v-bind('exThemeVars.sidebarColor');
        border-right: 1px solid v-bind('exThemeVars.splitColor');
    }
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.3s ease;
}
</style>
