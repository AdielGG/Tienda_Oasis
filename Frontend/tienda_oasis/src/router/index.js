import { createRouter, createWebHistory } from "vue-router";
import QuejasSugerenciasView from "@/views/QuejasSugerenciasView.vue";
import HomeView from "../views/HomeView.vue";
import LoginView from "@/views/login/LoginView.vue";
import RegisterView from "@/views/login/RegsterView.vue";
import AboutView from "@/views/AboutView.vue";
import ProductsView from "@/views/products/ProductsView.vue";
import AdminView from "@/views/admin/AdminView.vue";
import AdminUsersView from "@/views/admin/AdminUsersView.vue";
import AdminProductsView from "@/views/admin/AdminProductsView.vue";
import AddUser from "@/views/admin/AddUser.vue";
import AddProduct from "@/views/admin/AddProduct.vue";
import FaqView from "@/views/FaqView.vue";
import QuejasSugerencias from "@/components/QuejasSugerencias.vue";
import AdminSugerencias from "@/views/admin/AdminSugerencias.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/faq",
      name: "faq",
      component: FaqView,
    },{
      path: "/quejas-sugerencias",
      name: "quejas-sugerencias",
      component: QuejasSugerenciasView,
    },
    {
      path: "/about",
      name: "about",
      component: AboutView,
    },
    {
      path: "/register",
      name: "register",
      component: RegisterView,
    },
    {
      path: "/login",
      name: "login",
      component: LoginView,
    },
    {
      path: "/products",
      name: "products",
      component: ProductsView,
    },
    {
      path: "/admin",
      name: "admin",
      component: AdminView,
    },
    {
      path: "/admin/users",
      name: "users",
      component: AdminUsersView,
    },
    
    {
      path: "/admin/sugerencias",
      name: "sugerencias",
      component: AdminSugerencias,
    },
    {
      path: "/admin/users/add",
      name: "add-users",
      component: AddUser,
    },
    {
      path: "/admin/products",
      name: "admin-products",
      component: AdminProductsView,
    },
    {
      path: "/admin/products/add",
      name: "add-products",
      component: AddProduct,
    },
    
  ],
});

export default router;
