<script setup lang="ts">
    // biome-ignore assist/source/organizeImports: 誤判定
    import { RouterLink, RouterView } from "vue-router";
</script>

<template>
    <header>
        <nav>
            <span class="nav-title">Polyxia</span>
            <RouterLink class="item" exact-active-class="active" to="/">QuakeNow</RouterLink>
            <RouterLink id="nav-item-tsunami" class="item" exact-active-class="active" to="/tsunami">津波情報</RouterLink>
        </nav>
        <div class="header-right"><RouterLink class="item" exact-active-class="active" to="/about">About</RouterLink></div>
    </header>
    <RouterView id="view" />
</template>

<style scoped>
    header {
        display: flex;
        justify-content: space-between;
        position: fixed;
        z-index: 1000;
        top: 0;
        left: 0;
        width: 100vw;
        background: #fbfbfbb5;
        backdrop-filter: blur(18px);

        & nav {
            width: 100%;
            font-size: 13px;
            margin: 1rem 0 1rem 1rem;

            & .nav-title {
                font-size: 16px;
                font-weight: bold;
                margin-right: 32px;
            }

            & .item {
                margin-right: 8px;
                padding: 12px;
                transition: 0.2s;
                /* これがないとtransitionが不自然になる */
                border-bottom: solid transparent 1px;
            }

            & .active {
                color: black;
                font-weight: bold;
                border-bottom: solid black 1px;
            }
        }

        & .header-right {
            margin: 1rem;

            & .item {
                transition: 0.2s;
            }

            & .active {
                color: black;
                font-weight: bold;
            }
        }
    }

    #view {
        margin-top: 52px;
        width: 100%;
        height: 100%;
    }

    @keyframes pulse-glow-watch {
        0%, 100% {
            opacity: 0.6;
            text-shadow: 0 0 4px #8f0000;
        }
        50% {
            opacity: 1;
            text-shadow: 0 0 4px #8f0000, 0 0 12px #ff4d4d, 0 0 20px #ff4d4d;
        }
    }

    @keyframes pulse-glow-warning {
        0%, 100% {
            opacity: 0.6;
            text-shadow: 0 0 4px #e60000;
        }
        50% {
            opacity: 1;
            text-shadow: 0 0 4px #e60000, 0 0 12px #ff3333, 0 0 25px #ff3333;
        }
    }

    /* 注意報 (Watch) */
    .tsunami-watch {
        color: #8f0000;
        font-weight: bold;
        animation: pulse-glow-watch 3s ease-in-out infinite;
    }

    /* 警報 (Warning) */
    .tsunami-warning {
        color: #e60000;
        font-weight: bold;
        animation: pulse-glow-warning 2.2s ease-in-out infinite;
    }
</style>
