const storage_key = 'auth-user';

export type AuthUser = {
  id: string;
  email: string;
  name: string;
  session_token: string;
};

export const saveAuthUser = (authUser: AuthUser) => {
  localStorage.setItem(storage_key, JSON.stringify(authUser));
};

export const loadAuthUser = (): AuthUser | null => {
  const item = localStorage.getItem(storage_key);
  if (!item) {
    return null;
  }

  const user = JSON.parse(item);
  if (!user || typeof user !== 'object' || !('id' in user) || !('email' in user) || !('name' in user) || !('session_token' in user)) {
    return null;
  }

  return user;
};

export const removeAuthUser = () => {
  localStorage.removeItem(storage_key);
};
