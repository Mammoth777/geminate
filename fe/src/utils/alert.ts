import Swal from 'sweetalert2';
import { createApp, defineComponent, h } from 'vue';

// 成功提示
export const showSuccess = (message: string) => {
  return Swal.fire({
    title: '成功',
    text: message,
    icon: 'success',
    confirmButtonText: '确定',
    confirmButtonColor: '#52c41a',
    timer: 2000,
    timerProgressBar: true
  });
};

// 错误提示
export const showError = (message: string) => {
  return Swal.fire({
    title: '错误',
    text: message,
    icon: 'error',
    confirmButtonText: '确定',
    confirmButtonColor: '#f5222d'
  });
};

// 警告提示
export const showWarning = (message: string) => {
  return Swal.fire({
    title: '警告',
    text: message,
    icon: 'warning',
    confirmButtonText: '确定',
    confirmButtonColor: '#faad14'
  });
};

// 普通信息提示
export const showInfo = (message: string) => {
  return Swal.fire({
    title: '提示',
    text: message,
    icon: 'info',
    confirmButtonText: '确定',
    confirmButtonColor: '#1890ff'
  });
};

// 确认对话框
export const showConfirm = (title: string, text: string, confirmButtonText = '确定') => {
  return Swal.fire({
    title,
    text,
    icon: 'question',
    showCancelButton: true,
    confirmButtonColor: '#1890ff',
    cancelButtonColor: '#d9d9d9',
    confirmButtonText,
    cancelButtonText: '取消'
  });
};

// 
export const showDialog = (title: string, html: string, confirmButtonText = '确定') => {
  return Swal.fire({
    title,
    html,
    // showCancelButton: true,
    confirmButtonColor: '#1890ff',
    cancelButtonColor: '#d9d9d9',
    confirmButtonText,
  });
};

// 加载中
export const showLoading = (title = '处理中...') => {
  return Swal.fire({
    title,
    allowOutsideClick: false,
    didOpen: () => {
      Swal.showLoading();
    }
  });
};

// 关闭加载
export const closeLoading = () => {
  Swal.close();
};

// 带导航功能的对话框
export const showNavigateDialog = (title: string, text: string, url: string, buttonText = '前往查看') => {
  return Swal.fire({
    title,
    text,
    icon: 'info',
    showCancelButton: true,
    confirmButtonColor: '#1890ff',
    cancelButtonColor: '#d9d9d9',
    confirmButtonText: buttonText,
    cancelButtonText: '取消'
  }).then((result) => {
    if (result.isConfirmed) {
      window.location.href = url;
    }
    return result;
  });
};

// 带有 Vue 组件渲染能力的对话框
export const showComponentDialog = (title: string, component: any, props: any = {}, options: any = {}) => {
  // 创建挂载点
  const mountPoint = document.createElement('div');
  document.body.appendChild(mountPoint);

  // 创建包装组件
  const wrapper = defineComponent({
    render() {
      return h(component, props);
    }
  });

  // 创建应用并挂载
  const app = createApp(wrapper);
  app.mount(mountPoint);

  // 配置 SweetAlert2 选项
  const swalOptions = {
    title,
    html: mountPoint,
    showConfirmButton: true,
    confirmButtonText: options.confirmButtonText || '关闭',
    confirmButtonColor: options.confirmButtonColor || '#1890ff',
    width: options.width || '600px',
    ...options,
    didClose: () => {
      // 在对话框关闭时卸载组件并清理
      app.unmount();
      if (options.didClose) options.didClose();
    }
  };

  // 显示对话框
  return Swal.fire(swalOptions);
};
