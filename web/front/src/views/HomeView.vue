<template>
  <el-container class="contain">
    <el-header class="containHeader">
      <TopBar />
    </el-header>
    <el-main class="containMain">
      <el-row :gutter="20">
        <el-col :span="4">
          <el-affix :offset="71">
            <div class="navCard" >
              <NavCard/>
            </div>
          </el-affix>
        </el-col>
        <el-col :span="16" justify="center">

          <div class="articleCard">
            <router-view/>
          </div>
        </el-col>
        <el-col :span="4">
          <el-affix :offset="71">
            <div class="CommentCard">
              <CommentCard/>
            </div>
          </el-affix>

        </el-col>
      </el-row>
    </el-main>
    <el-footer class="containFooter">
      <FooterBar />
    </el-footer>
  </el-container>
</template>

<script lang="ts" setup>
import TopBar from "@/components/TopBar.vue";
import NavCard from "@/components/homePage/NavCard.vue";
import CommentCard from "@/components/homePage/CommentCard.vue";
import FooterBar from "@/components/FooterBar.vue";
import { onMounted } from "vue";

onMounted(() => {});

</script>

<style scoped>
.contain {
    min-height: 100%;
    display: flex;
    flex-direction: column;
}

.containHeader {
}

.containMain {
    flex: 1;
}

.containFooter {
    align-items: center;
    flex-shrink: 0;
}

</style>
