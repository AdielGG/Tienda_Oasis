import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import LoginView from "@/views/login/LoginView.vue";
import RegisterView from "@/views/login/RegsterView.vue";
import AboutView from "@/views/AboutView.vue";
import ProductsView from "@/views/products/ProductsView.vue";
import AdminView from "@/views/admin/AdminView.vue";
import AdminUsersView from "@/views/admin/AdminUsersView.vue";
import AdminProductsView from "@/views/admin/AdminProductsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
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
      path: "/admin/products",
      name: "admin-products",
      component: AdminProductsView,
    },
  ],
});

export default router;
