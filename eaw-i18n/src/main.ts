import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';

// Global style
import '@/assets/styles/main.scss';

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
