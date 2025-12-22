<template>
  <div class="max-w-4xl mx-auto space-y-10 pb-24 md:pb-12">
    <!-- Premium Hero Header -->
    <div class="relative overflow-hidden rounded-[2.5rem] md:rounded-[3.5rem] bg-gradient-to-br from-slate-800 to-slate-950 text-white p-6 md:p-16 shadow-2xl">
      <div class="absolute top-0 right-0 -mt-20 -mr-20 w-96 h-96 bg-indigo-600/20 rounded-full blur-[100px]"></div>
      <div class="absolute bottom-0 left-0 -mb-20 -ml-20 w-80 h-80 bg-rose-600/10 rounded-full blur-[80px]"></div>
      
      <div class="relative z-10 flex flex-col md:flex-row items-center gap-8 md:gap-12">
        <div class="relative group">
          <div class="w-28 h-28 md:w-40 md:h-40 rounded-[2.5rem] bg-white/5 backdrop-blur-2xl border-4 border-white/10 overflow-hidden shadow-2xl group-hover:scale-105 transition-transform duration-500">
            <img v-if="user?.photo_url" :src="user.photo_url" alt="Profile" class="w-full h-full object-cover" />
            <div v-else class="w-full h-full flex items-center justify-center text-5xl font-black text-white/20">
              {{ user?.first_name?.charAt(0) || "U" }}
            </div>
          </div>
          <button class="absolute -bottom-2 -right-2 bg-indigo-600 text-white p-3 rounded-2xl shadow-xl hover:bg-indigo-700 transition-all hover:scale-110 active:scale-95 border-4 border-slate-900">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"></path></svg>
          </button>
        </div>
        
        <div class="text-center md:text-left space-y-4">
          <div class="space-y-1">
            <h1 class="text-3xl md:text-6xl font-black tracking-tighter leading-none">{{ user?.first_name }} {{ user?.last_name }}</h1>
            <p class="text-slate-400 text-lg font-medium">@{{ user?.username || "explorer" }}</p>
          </div>
          <div class="flex flex-wrap justify-center md:justify-start gap-3">
            <span v-if="user?.is_premium" class="px-4 py-1.5 bg-indigo-500/20 text-indigo-300 text-[10px] font-black rounded-full uppercase tracking-[0.2em] border border-indigo-500/30 backdrop-blur-xl">Premium Member</span>
            <span v-if="user?.is_verified" class="px-4 py-1.5 bg-emerald-500/20 text-emerald-300 text-[10px] font-black rounded-full uppercase tracking-[0.2em] border border-emerald-500/30 backdrop-blur-xl">Verified Account</span>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <!-- Preferences Section -->
      <section class="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 p-6 md:p-10 shadow-sm space-y-8">
        <div class="flex items-center gap-4">
          <div class="w-12 h-12 rounded-2xl bg-indigo-100 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400 flex items-center justify-center shadow-sm">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"></path></svg>
          </div>
          <div>
            <h2 class="text-2xl font-black text-slate-900 dark:text-white tracking-tight">Preferences</h2>
            <p class="text-sm text-slate-500 font-medium">Customize your experience.</p>
          </div>
        </div>

        <div class="space-y-6">
          <div v-for="setting in settings" :key="setting.id" class="flex items-center justify-between p-4 rounded-3xl hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-all duration-300 group">
            <div class="flex items-center gap-4">
              <div :class="setting.bg" class="w-10 h-10 md:w-12 md:h-12 rounded-2xl flex items-center justify-center group-hover:scale-110 transition-transform">
                <component :is="setting.icon" class="w-6 h-6" :class="setting.color" />
              </div>
              <div>
                <h4 class="text-base font-bold text-slate-900 dark:text-white">{{ setting.name }}</h4>
                <p class="text-xs text-slate-500 font-medium">{{ setting.description }}</p>
              </div>
            </div>
            
            <div v-if="setting.name === 'Dark Mode'" class="flex items-center bg-slate-100 dark:bg-slate-800 p-1 rounded-xl">
              <button v-for="theme in ['light', 'dark', 'system']" :key="theme" @click="authStore.setTheme(theme)" class="px-3 py-1.5 rounded-lg text-[10px] font-black uppercase tracking-widest transition-all" :class="authStore.theme === theme ? 'bg-white dark:bg-slate-700 text-indigo-600 dark:text-indigo-400 shadow-sm' : 'text-slate-400 hover:text-slate-600 dark:hover:text-slate-200'">
                {{ theme }}
              </button>
            </div>
            
            <button v-else @click="setting.enabled = !setting.enabled" class="w-14 h-8 rounded-full transition-all relative p-1" :class="setting.enabled ? 'bg-indigo-600' : 'bg-slate-200 dark:bg-slate-700'">
              <div class="w-6 h-6 bg-white rounded-full shadow-lg transition-transform duration-300" :class="setting.enabled ? 'translate-x-6' : 'translate-x-0'"></div>
            </button>
          </div>
        </div>
      </section>

      <!-- Integrations & Account Section -->
      <div class="space-y-8">
        <section class="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 p-6 md:p-10 shadow-sm space-y-8">
          <div class="flex items-center gap-4">
            <div class="w-12 h-12 rounded-2xl bg-blue-100 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400 flex items-center justify-center shadow-sm">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path></svg>
            </div>
            <div>
              <h2 class="text-2xl font-black text-slate-900 dark:text-white tracking-tight">Integrations</h2>
              <p class="text-sm text-slate-500 font-medium">Connect with other tools.</p>
            </div>
          </div>

          <button class="w-full flex items-center justify-between p-4 md:p-6 rounded-[2rem] border-2 border-slate-50 dark:border-slate-800 hover:border-indigo-100 dark:hover:border-indigo-900/50 hover:bg-indigo-50/30 dark:hover:bg-indigo-900/10 transition-all duration-500 group">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 bg-blue-500 text-white rounded-2xl flex items-center justify-center shadow-lg shadow-blue-200 dark:shadow-none group-hover:scale-110 transition-transform">
                <svg class="w-7 h-7" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm4.64 6.8c-.15 1.58-.8 5.42-1.13 7.19-.14.75-.42 1-.68 1.03-.58.05-1.02-.38-1.58-.75-.88-.58-1.38-.94-2.23-1.5-.99-.65-.35-1.01.22-1.59.15-.15 2.71-2.48 2.76-2.69.01-.03.01-.14-.07-.2-.08-.06-.19-.04-.27-.02-.12.02-1.96 1.25-5.54 3.67-.52.36-.99.53-1.42.52-.47-.01-1.37-.26-2.03-.48-.82-.27-1.47-.42-1.42-.88.03-.24.37-.48 1.02-.73 3.99-1.73 6.64-2.87 7.95-3.42 3.78-1.58 4.57-1.85 5.08-1.86.11 0 .37.03.53.17.14.11.18.26.2.37.01.06.03.24.01.37z"/></svg>
              </div>
              <div class="text-left">
                <h4 class="text-base font-bold text-slate-900 dark:text-white">Telegram Bot</h4>
                <p class="text-xs text-slate-500 font-medium">Manage tasks via Telegram</p>
              </div>
            </div>
            <div class="px-4 py-1.5 rounded-xl text-[10px] font-black uppercase tracking-widest" :class="user?.telegram_id !== 0 ? 'bg-emerald-500/10 text-emerald-500' : 'bg-slate-100 dark:bg-slate-800 text-slate-400'">
              {{ user?.telegram_id !== 0 ? 'Connected' : 'Connect' }}
            </div>
          </button>
        </section>

        <section class="bg-white dark:bg-slate-900 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 p-8 md:p-10 shadow-sm space-y-8">
          <button @click="handleLogout" class="w-full py-5 px-8 bg-rose-500 text-white rounded-[2rem] font-black text-xs uppercase tracking-[0.2em] hover:bg-rose-600 hover:scale-[1.02] active:scale-[0.98] transition-all shadow-xl shadow-rose-200 dark:shadow-none flex items-center justify-center gap-3">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path></svg>
            <span>Terminate Session</span>
          </button>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, defineComponent, h } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";

const router = useRouter();
const authStore = useAuthStore();
const user = computed(() => authStore.user);

const IconBell = defineComponent({
  render: () =>
    h(
      "svg",
      {
        fill: "none",
        stroke: "currentColor",
        viewBox: "0 0 24 24",
        class: "w-inherit h-inherit",
      },
      [
        h("path", {
          "stroke-linecap": "round",
          "stroke-linejoin": "round",
          "stroke-width": "2",
          d: "M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9",
        }),
      ]
    ),
});

const IconMoon = defineComponent({
  render: () =>
    h(
      "svg",
      {
        fill: "none",
        stroke: "currentColor",
        viewBox: "0 0 24 24",
        class: "w-inherit h-inherit",
      },
      [
        h("path", {
          "stroke-linecap": "round",
          "stroke-linejoin": "round",
          "stroke-width": "2",
          d: "M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z",
        }),
      ]
    ),
});

const IconLock = defineComponent({
  render: () =>
    h(
      "svg",
      {
        fill: "none",
        stroke: "currentColor",
        viewBox: "0 0 24 24",
        class: "w-inherit h-inherit",
      },
      [
        h("path", {
          "stroke-linecap": "round",
          "stroke-linejoin": "round",
          "stroke-width": "2",
          d: "M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z",
        }),
      ]
    ),
});

const settings = ref([
  {
    id: 1,
    name: "Push Notifications",
    description: "Receive alerts for upcoming tasks and goal progress.",
    icon: IconBell,
    bg: "bg-indigo-50",
    color: "text-indigo-600",
    enabled: true,
  },
  {
    id: 2,
    name: "Dark Mode",
    description: "Switch to a darker theme for better night viewing.",
    icon: IconMoon,
    bg: "bg-slate-900",
    color: "text-white",
    enabled: false,
  },
  {
    id: 3,
    name: "Privacy Mode",
    description: "Hide sensitive task details in the dashboard view.",
    icon: IconLock,
    bg: "bg-rose-50",
    color: "text-rose-600",
    enabled: false,
  },
]);

const handleLogout = () => {
  authStore.logout();
  router.push("/login");
};
</script>
