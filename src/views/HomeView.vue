<script setup lang="ts">
import { ref, onMounted } from "vue";

import QuakeMap from "@/components/QuakeMap.vue";

const eventId = ref("");
const isDebug = ref(false);
const isLoading = ref(true);

const fetchEventId = async (index: number): Promise<string> => {
    const resp = await fetch("https://quake-jade.vercel.app/api/events");
    const events = (await resp.json() as { events: string[] }).events;

    // IDに_VXSE51または_VXSE53が含まれるもののうちindex番目のイベントを取得
    const filtered = events.filter((id) => id.includes("_VXSE51") || id.includes("_VXSE53"));
    if (index < 0 || index >= filtered.length) {
        throw new Error("Index out of bounds");
    }

    return filtered[index];
};

onMounted(async () => {
    if (window.localStorage.getItem("DBG_DUMMY_EVENT_ID")) {
        isDebug.value = true;
        eventId.value = window.localStorage.getItem("DBG_DUMMY_EVENT_ID")!;
        isLoading.value = false;

        return;
    }

    // 最初のイベントを取得
    eventId.value = await fetchEventId(0);

    isLoading.value = false;
});
</script>

<template>
    <main>
        <QuakeMap v-if=!isLoading :is-debug="isDebug" :event-id="eventId" />
    </main>
</template>
