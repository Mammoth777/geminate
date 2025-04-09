import axios from 'axios';

// 定义工作流接口
export interface Workflow {
  id: string;
  name: string;
  description: string;
  status: string;
  createAt: string;
  updateAt: string;
  nodes: any[];
  edges: any[];
}

// 工作流服务
export const workflowService = {
  // 获取工作流列表
  getWorkflows: async (params?: { page?: number; pageSize?: number }) => {
    const queryParams = new URLSearchParams();
    
    if (params?.page) {
      queryParams.append('page', params.page.toString());
    }
    
    if (params?.pageSize) {
      queryParams.append('pageSize', params.pageSize.toString());
    }
    
    const queryString = queryParams.toString();
    const url = `/api/workflows${queryString ? '?' + queryString : ''}`;
    
    const response = await axios.get(url);
    return response.data;
  },
  
  // 获取单个工作流详情
  getWorkflow: async (id: string) => {
    const response = await axios.get(`/api/workflows/${id}`);
    return response.data;
  },
  
  // 保存工作流
  saveWorkflow: async (data: any) => {
    if (data.id) {
      // 更新现有工作流
      const response = await axios.put(`/api/workflows/${data.id}`, data);
      return response.data;
    } else {
      // 创建新工作流
      const response = await axios.post('/api/workflows', data);
      return response.data;
    }
  },
  
  // 删除工作流
  deleteWorkflow: async (id: string) => {
    const response = await axios.delete(`/api/workflows/${id}`);
    return response.data;
  },

  publishWorkflow: async (id: string) => {
    const response = await axios.post(`/api/workflows/${id}/publish`);
    return response.data;
  }
};