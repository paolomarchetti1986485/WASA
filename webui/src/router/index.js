import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'

const routes = [
  { path: '/login', component: LoginView },
  { 
    path: '/', 
    component: HomeView,
    meta: { requiresAuth: true } // Aggiungi meta per richiedere l'autenticazione
  },
  { path: '/link1', component: HomeView, meta: { requiresAuth: true } },
  { path: '/link2', component: HomeView, meta: { requiresAuth: true } },
  { path: '/some/:id/link', component: HomeView, meta: { requiresAuth: true } },
  { path: '/:pathMatch(.*)*', redirect: '/login' } // Redirezione a login per qualsiasi rotta non trovata
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const loggedIn = localStorage.getItem('token')

  if (to.matched.some(record => record.meta.requiresAuth) && !loggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
