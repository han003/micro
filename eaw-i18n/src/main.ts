import { MdButton } from 'vue-material/dist/components';
import 'vue-material/dist/vue-material.min.css';
import 'vue-material/dist/theme/default.css';

import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';

// Global style
import '@/assets/styles/main.scss';

Vue.use(MdButton);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
