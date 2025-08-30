import { atom } from 'jotai'

import { models } from '@wailsjs/go/models'

export const apps_catalog_atom = atom<models.AppCatalog[]>([
  {
    name: 'Bazel',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  },
  {
    name: 'Bazel',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  },
  {
    name: 'Bazel',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  },
  {
    name: 'Bazel',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  }
])
export const websites_catalog_atom = atom<models.WebsiteCatalog[]>([
  {
    name: 'Bazel',
    url: 'www.example.com',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  },
  {
    name: 'Bazel',
    url: 'www.example.com',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  },
  {
    name: 'Bazel',
    url: 'www.example.com',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  },
  {
    name: 'Bazel',
    url: 'www.example.com',
    version: '1.1.1',
    image_url:
      'https://imgs.search.brave.com/pJm6U5isygwaXJ35h_Kwrn0hRBKTo0KHLS4XsbvGhkY/rs:fit:500:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YTIuZGV2LnRvL2R5/bmFtaWMvaW1hZ2Uv/d2lkdGg9MTkwLGhl/aWdodD0sZml0PXNj/YWxlLWRvd24sZ3Jh/dml0eT1hdXRvLGZv/cm1hdD1hdXRvL2h0/dHBzOi8vZGV2LXRv/LXVwbG9hZHMuczMu/YW1hem9uYXdzLmNv/bS91cGxvYWRzL2Fy/dGljbGVzLzhqN2t2/cDY2MHJxenQ5OXp1/aThlLnBuZw'
  }
])
