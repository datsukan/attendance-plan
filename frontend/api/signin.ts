import axios from 'axios';

type result = {
  id: string;
  email: string;
  name: string;
  createdAt: string;
  updatedAt: string;
  sessionToken: string;
};

export const signin = async (email: string, password: string): Promise<result | null> => {
  const param = { email, password };

  try {
    const response = await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/signin`, param, {
      headers: {
        'Content-Type': 'application/json',
      },
    });

    const result: result = {
      id: response.data.id,
      email: response.data.email,
      name: response.data.name,
      createdAt: response.data.created_at,
      updatedAt: response.data.updated_at,
      sessionToken: response.data.session_token,
    };

    return result;
  } catch (error) {
    console.error(error);
  }

  return null;
};
