import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '@/views/ProfileView.vue'
import NonLoggedProfileView from '@/views/NonLoggedProfileView.vue'
import UploadPhoto from '@/views/UploadPhotoView.vue'; 

const routes = [
  { 
    path: '/login', 
    component: LoginView 
  },
  { 
    path: '/', 
    component: HomeView,
    meta: { requiresAuth: true } // Richiede autenticazione
  },
  { 
    path: '/user/:userId/profile',
    name: 'Profile',
    component: ProfileView,
    props: true,
    meta: { requiresAuth: true } // Richiede autenticazione
  },
  { 
    path: '/user/:userId/nonlogged-profile', 
    name: 'NonLoggedProfile', 
    component: NonLoggedProfileView,
    props: true,
    meta: { requiresAuth: true } // Richiede autenticazione
  },
  {
    path: '/search',
    name: 'SearchUser',
    component: () => import('../views/SearchView.vue'),  // Assumi che esista una view dedicata alla ricerca
    meta: { requiresAuth: true } // Richiede autenticazione
  },
  {
    path: '/upload',
    name: 'UploadPhoto',
    component: UploadPhoto,
    meta: { requiresAuth: true } // Richiede autenticazione
  },
  // Route catch-all per non trovare percorsi validi
  { 
    path: '/:pathMatch(.*)*', 
    redirect: '/login' 
  }
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  console.log("Token verificato in beforeEach:", token);

  // Se il token è nullo o non valido, reindirizza al login
  if (to.matched.some(record => record.meta.requiresAuth) && (!token || token === 'null' || token === 'undefined')) {
    console.log("Token non valido o mancante, reindirizzamento a /login");
    next('/login');
  } else {
    console.log("Token valido, proseguo alla rotta:", to.path);
    next();
  }
});


export default router
