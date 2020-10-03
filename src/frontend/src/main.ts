import { createApp } from 'vue'

import App from './App.vue'
import store from './store/index'

import './assets/styles/index.css'

// this adds the fun D&D font
require('typeface-libre-baskerville')

createApp(App).use(store).mount('#app')
