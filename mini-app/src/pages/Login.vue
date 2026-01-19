<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-50 p-4">
    <div class="w-full max-w-md">
      <!-- Logo & Header -->
      <div class="text-center mb-10">
        <div
          class="inline-flex items-center justify-center w-20 h-20 bg-indigo-600 rounded-3xl shadow-xl shadow-indigo-200 mb-6 transform -rotate-6"
        >
          <svg
            class="w-10 h-10 text-white"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
            />
          </svg>
        </div>
        <h1 class="text-4xl font-black text-slate-900 tracking-tight">
          Taskly
        </h1>
        <p class="text-slate-500 mt-2 font-medium">
          Your personal productivity hub
        </p>
      </div>

      <!-- Main Card -->
      <div
        class="bg-white rounded-[2.5rem] shadow-2xl shadow-slate-200/60 p-8 md:p-10 border border-slate-100"
      >
        <div v-if="loading" class="text-center py-10">
          <div
            class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-slate-100 border-t-indigo-600 mb-4"
          ></div>
          <p class="text-slate-600 font-bold">Verifying account...</p>
        </div>

        <div v-else-if="error" class="text-center py-6">
          <div
            class="w-16 h-16 bg-rose-50 text-rose-500 rounded-full flex items-center justify-center mx-auto mb-4"
          >
            <svg
              class="w-8 h-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
              />
            </svg>
          </div>
          <h2 class="text-xl font-bold text-slate-900 mb-2">
            Authentication Error
          </h2>
          <p class="text-slate-500 mb-8 text-sm leading-relaxed">{{ error }}</p>

          <div class="space-y-3">
            <button
              @click="handleLogin"
              class="w-full bg-indigo-600 text-white py-4 rounded-2xl font-bold hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100"
            >
              Try Again
            </button>

          </div>
        </div>

        <div v-else class="text-center">
          <div class="mb-8">
            <h2 class="text-2xl font-bold text-slate-900 mb-2">Welcome!</h2>
            <p class="text-slate-500 text-sm">
              Please sign in with Telegram to continue to your workspace.
            </p>
          </div>

          <div class="space-y-4">
            <!-- Telegram Widget Container -->
            <div id="telegram-widget-container" class="flex justify-center my-4"></div>


          </div>

          <p
            class="mt-8 text-[10px] font-bold text-slate-400 uppercase tracking-widest"
          >
            🔒 Secure & Private
          </p>
        </div>
      </div>

      <!-- Footer Info -->
      <div class="mt-10 grid grid-cols-3 gap-4">
        <div class="text-center">
          <div class="text-xl mb-1">⚡</div>
          <p
            class="text-[10px] font-bold text-slate-500 uppercase tracking-tighter"
          >
            Fast
          </p>
        </div>
        <div class="text-center">
          <div class="text-xl mb-1">🛡️</div>
          <p
            class="text-[10px] font-bold text-slate-500 uppercase tracking-tighter"
          >
            Secure
          </p>
        </div>
        <div class="text-center">
          <div class="text-xl mb-1">☁️</div>
          <p
            class="text-[10px] font-bold text-slate-500 uppercase tracking-tighter"
          >
            Sync
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";
import api from "../services/api";

const router = useRouter();
const authStore = useAuthStore();
const loading = ref(false);
const error = ref(null);

async function handleLogin() {
  loading.value = true;
  error.value = null;

  try {
    if (window.Telegram?.WebApp?.initData) {
      await authStore.loginWithTelegram(window.Telegram.WebApp.initData);
      router.push("/");
    } else {
      // If not in Telegram WebApp, show error but widget is also available
      throw new Error(
        "No Telegram data available. Please open this app from Telegram or use the login widget."
      );
    }
  } catch (err) {
    error.value =
      err.response?.data?.error || err.message || "Authentication failed";
    console.error("Login error:", err);
  } finally {
    loading.value = false;
  }
}



// Global callback for Telegram Widget
window.onTelegramAuth = async (user) => {
  loading.value = true;
  error.value = null;
  try {
    await authStore.loginWithWidget(user);
    router.push("/");
  } catch (err) {
    error.value = "Widget login failed";
    console.error(err);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  if (window.Telegram?.WebApp) {
    window.Telegram.WebApp.ready();
    if (window.Telegram.WebApp.initData) {
      handleLogin();
    }
  } else if (!authStore.isAuthenticated) {
    // Automatically enter demo mode if not in Telegram and not authenticated
    // enterPreviewMode(); 
    // Commented out auto-demo to allow widget usage
  }

  // Inject Telegram Widget script dynamically
  const script = document.createElement("script");
  script.async = true;
  script.src = "https://telegram.org/js/telegram-widget.js?22";
  script.setAttribute("data-telegram-login", "OptiList_bot");
  script.setAttribute("data-size", "large");
  script.setAttribute("data-onauth", "onTelegramAuth(user)");
  script.setAttribute("data-request-access", "write");
  document.getElementById("telegram-widget-container").appendChild(script);
});
</script>
