<template>
    <div id="app">
        <!-- 页眉 -->
        <header>
            <h1>PDF Online</h1>
        </header>

        <div class="maintainer">
            <div class="left">
                <div class="theme">主题</div>
                <div class="login">登录</div>
                <div class="upload">上传</div>
                <div class="home">首页</div>
            </div>
            <div class="right">
                <!-- PDF列表 -->
                <div class="pdf-container">
                    <div v-for="pdf in pdfList" :key="pdf.title" class="pdf-block" @click="openPDF(pdf.url)">
                        <!-- <img :src="pdf.cover_url" alt="加载封面失败" class="cover"> -->
                        <el-image
                            style="width: 100px; height: 100px"
                            :src="pdf.cover_url"
                            :fit="fit"></el-image>
                        <p class="title">{{ pdf.title }}</p>
                        <p class="desc">{{ pdf.description }}</p>
                    </div>
                </div>
            </div>
        </div>

        <!-- 分页控制 -->
        <!-- <div class="pagination">
            <button @click="prevPage" :disabled="currentPage === 1">
                上一页
            </button>
            <span>{{ currentPage }} / {{ totalPages }}</span>
            <button @click="nextPage" :disabled="currentPage === totalPages">
                下一页
            </button>
        </div> -->

        <!-- 底部功能按钮 -->
        <!-- <footer>
            <button @click="showMore">更多</button>
            <button @click="login">登录</button>
            <button @click="goHome">首页</button>
            <button @click="toggleBackgroundMode">背景模式</button>
            <button @click="upload">上传</button>
        </footer> -->
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            pdfList: [], // 所有 PDF 数据
            displayedPDFs: [], // 当前页面显示的 PDF 数据
            itemsPerPage: 8, // 每页显示的 PDF 数量
            currentPage: 1, // 当前页数
        };
    },
    created() {
        this.fetchPDFList();
    },
    methods: {
        async fetchPDFList() {
            try {
                const response = await this.$axios.get(
                    "http://localhost:8080/v1/pdfs"
                );
                this.pdfList = response.data.pdfs;
                this.downloadBaseURL = "http://localhost:8080";
                console.log("pdfList:", this.pdfList);
            } catch (error) {
                console.error("Error fetching PDF list:", error);
            }
        },
        openPDF(pdfFileName) {
            console.log("filename:" + pdfFileName);
            // const fileName = pdfFileName.substring(
            //     pdfFileName.lastIndexOf("/") + 1
            // );
            const pdfUrl = this.downloadBaseURL + pdfFileName;
            window.location.href = pdfUrl;
        },
        prevPage() {
            if (this.currentPage > 1) {
                this.currentPage--;
                this.updateDisplayedPDFs();
            }
        },
        nextPage() {
            if (this.currentPage < this.totalPages) {
                this.currentPage++;
                this.updateDisplayedPDFs();
            }
        },
        updateDisplayedPDFs() {
            const startIdx = (this.currentPage - 1) * this.itemsPerPage;
            const endIdx = startIdx + this.itemsPerPage;
            this.displayedPDFs = this.pdfList.slice(startIdx, endIdx);
        },
        showMore() {
            // 处理“更多”按钮点击事件
            // ...
            alert("more");
        },
        login() {
            // 处理“登录”按钮点击事件
            // ...
            alert("login");
        },
        goHome() {
            // 处理“首页”按钮点击事件
            // ...
            alert("gohome");
        },
        toggleBackgroundMode() {
            // 处理“背景模式”按钮点击事件
            // ...
            alert("backgroud");
        },
        upload() {
            // 处理“上传”按钮点击事件
            // ...
            alert("upload");
        },
    },
    computed: {
        totalPages() {
            return Math.ceil(this.pdfList.length / this.itemsPerPage);
        },
    },
};
</script>
  
<style>
@import "css/index.css";
</style>

  