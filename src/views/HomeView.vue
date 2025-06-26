<script setup lang="ts">
import { ref, onMounted } from "vue";

import QuakeMap from "@/components/QuakeMap.vue";

const eventId = ref("");
const isDebug = ref(false);
const isLoading = ref(true);
const currentEventIndex = ref(0);
const isSeekableToNext = ref(false);
const isSeekableToPrev = ref(false);

let events: string[] | undefined = undefined;

const fetchEventId = async (index: number): Promise<{id: string, isLatest: boolean, isFinal: boolean}> => {
    if (!events) {
        throw new Error("Events not loaded");
    }

    // IDに_VXSE51または_VXSE53が含まれるもののうちindex番目のイベントを取得
    const filtered = events.filter((id) => id.includes("_VXSE51") || id.includes("_VXSE53"));
    if (index < 0 || index >= filtered.length) {
        throw new Error("Index out of bounds");
    }

    return { id: filtered[index], isLatest: index === 0, isFinal: index === filtered.length - 1 };
};

const updateEventId = async (index: number): Promise<void> => {
    console.log(`Updating event ID to index: ${index}`);
    const fetched = await fetchEventId(index);
    eventId.value = fetched.id;
    isSeekableToNext.value = !fetched.isLatest;
    isSeekableToPrev.value = !fetched.isFinal;
};

onMounted(async () => {
    if (window.localStorage.getItem("DBG_DUMMY_EVENT_ID")) {
        isDebug.value = true;
        eventId.value = window.localStorage.getItem("DBG_DUMMY_EVENT_ID")!;
        isLoading.value = false;

        return;
    }

    const resp = await fetch("https://quake-jade.vercel.app/api/events");
    events = (await resp.json() as { events: string[] }).events;

    // 最初のイベントを取得
    await updateEventId(currentEventIndex.value);

    isLoading.value = false;
});
</script>

<template>
    <main>
        <QuakeMap v-if=!isLoading :is-debug="isDebug" :event-id="eventId" />
        <div>
            <div v-if="isLoading" class="loading">Loading...</div>
            <div v-else class="controls">
                <button
                    :disabled="!isSeekableToNext"
                    @click="async () => {
                        currentEventIndex--;
                        await updateEventId(currentEventIndex);
                    }"
                >
                    Next
                </button>
                <button
                    :disabled="!isSeekableToPrev"
                    @click="async () => {
                        currentEventIndex++;
                        await updateEventId(currentEventIndex);
                    }"
                >
                    Previous
                </button>
            </div>
        </div>
    </main>
</template>
