<template>
    <el-container class="contain">
        <el-aside width="200px">Aside</el-aside>
        <el-container>
            <el-header>
                <Nav/>
            </el-header>
            <el-main>Main</el-main>
            <el-footer>
                <Footer/>
            </el-footer>
        </el-container>
    </el-container>
</template>
<script lang="ts" setup>
import Footer from "@/components/admin/Footer.vue"
import Nav from "@/components/admin/Nav.vue"


</script>
<style scoped>
.contain {
    height: 100%;
    width: auto;
}
</style>
