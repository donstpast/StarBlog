<template>
    <el-container class="contain">
        <el-aside width="230px" class="containAside">
            <Nav />
        </el-aside>
        <el-container>
            <el-header class="containHeader">
                <Header/>
            </el-header>
            <el-main class="containMain">
                <router-view></router-view>
            </el-main>
            <el-footer class="containFooter">
                <Footer/>
            </el-footer>
        </el-container>
    </el-container>
</template>
<script lang="ts" setup>
import Header  from "@/components/admin/Header.vue";
import Footer from "@/components/admin/Footer.vue"
import Nav from "@/components/admin/Nav.vue"
</script>
<style scoped>
.contain {
    height: 100%;
}
.containAside {
    background: rgb(51,58,69);
}
.containHeader {
    background: rgba(255,255,255,0.5);
    display: flex;
    justify-content: flex-end;
    align-items: center;
}
.containMain {
}
.containFooter {
    background: rgba(255,255,255,0.5);
    align-items: center;
}
</style>
