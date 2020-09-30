import { createApp } from 'vue'

import App from './App.vue'
import store from './store'

import './registerServiceWorker'
import './assets/styles/index.css'

// this adds the fun D&D font
require('typeface-libre-baskerville')

createApp(App).use(store).mount('#app')
