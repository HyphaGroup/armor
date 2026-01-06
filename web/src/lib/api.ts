const API_BASE = '/api';

function getPassword(): string | null {
  if (typeof window === 'undefined') return null;
  return sessionStorage.getItem('armor_password');
}

export function setPassword(password: string): void {
  sessionStorage.setItem('armor_password', password);
}

export function clearPassword(): void {
  sessionStorage.removeItem('armor_password');
}

export function hasPassword(): boolean {
  return !!getPassword();
}

async function request<T>(path: string, options: RequestInit = {}): Promise<T> {
  const password = getPassword();
  if (!password) {
    throw new Error('Not authenticated');
  }

  const response = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${password}`,
      ...options.headers,
    },
  });

  if (response.status === 401) {
    clearPassword();
    throw new Error('Invalid password');
  }

  if (!response.ok) {
    const error = await response.json().catch(() => ({ error: response.statusText }));
    throw new Error(error.error || response.statusText);
  }

  if (response.status === 204) {
    return undefined as T;
  }

  return response.json();
}

export interface ProfileSummary {
  id: string;
  name: string;
  description: string;
  completeness: number;
  created_at: string;
  updated_at: string;
}

export interface SectionCompleteness {
  section: string;
  percentage: number;
  filled: number;
  total: number;
}

export interface ProfileCompleteness {
  overall: number;
  sections: SectionCompleteness[];
}

export interface Profile {
  id: string;
  name: string;
  description: string;
  mission: any;
  assets: any;
  adversaries: any;
  threats: any;
  risks: any;
  mitigations: any;
  completeness: ProfileCompleteness;
  created_at: string;
  updated_at: string;
}

export interface ValidationError {
  path: string;
  message: string;
}

export const api = {
  async listProfiles(): Promise<ProfileSummary[]> {
    return request<ProfileSummary[]>('/profiles');
  },

  async createProfile(name: string, description: string = ''): Promise<Profile> {
    return request<Profile>('/profiles', {
      method: 'POST',
      body: JSON.stringify({ name, description }),
    });
  },

  async getProfile(id: string): Promise<Profile> {
    return request<Profile>(`/profiles/${id}`);
  },

  async deleteProfile(id: string): Promise<void> {
    return request<void>(`/profiles/${id}`, { method: 'DELETE' });
  },

  async getSection(profileId: string, section: string): Promise<{ data: any }> {
    return request<{ data: any }>(`/profiles/${profileId}/${section}`);
  },

  async updateSection(profileId: string, section: string, data: any): Promise<{ success: boolean; data: any }> {
    return request<{ success: boolean; data: any }>(`/profiles/${profileId}/${section}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  },

  async checkPassword(password: string): Promise<boolean> {
    try {
      const response = await fetch(`${API_BASE}/profiles`, {
        headers: {
          'Authorization': `Bearer ${password}`,
        },
      });
      return response.ok;
    } catch {
      return false;
    }
  },
};
