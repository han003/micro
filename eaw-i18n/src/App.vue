<template>
  <div class="page-container md-layout-row">
    <md-app id="page-app">
      <md-app-toolbar class="md-primary md-dense">
        <span class="md-title">Vendor management</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="clipped">
        <md-list>
          <md-list-item to="/">
            <md-icon>home</md-icon>
            <span class="md-list-item-text">Home</span>
          </md-list-item>

          <md-list-item to="/vendors">
            <md-icon>business</md-icon>
            <span class="md-list-item-text">Vendors</span>
          </md-list-item>

          <md-list-item to="/settings">
            <md-icon>settings</md-icon>
            <span class="md-list-item-text">Settings</span>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <router-view />
        <!--      <router-link to="/about">About</router-link>-->
      </md-app-content>
    </md-app>
  </div>
</template>
