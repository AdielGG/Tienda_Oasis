import { z } from 'zod';

export const loginSchema = z.object({
  username: z.string()
    .min(4, { message: 'El usuario debe tener al menos 4 caracteres' })
    .max(20, { message: 'El usuario no puede tener más de 20 caracteres' })
    .regex(/^[a-zA-Z]+$/, { message: 'El usuario solo puede contener letras' }),
  password: z.string()
    .min(4, { message: 'La contraseña debe tener al menos 4 caracteres' })
    .max(20, { message: 'La contraseña no puede tener más de 20 caracteres' })
});


export const registerSchema = z.object({
  name: z.string()
    .min(2, { message: 'El nombre debe tener al menos 2 caracteres' })
    .max(50, { message: 'El nombre no puede tener más de 50 caracteres' }),
  
  lastname: z.string()
    .min(2, { message: 'El apellido debe tener al menos 2 caracteres' })
    .max(50, { message: 'El apellido no puede tener más de 50 caracteres' }),
  
  age: z.number()
    .min(18, { message: 'Debes ser mayor de edad (18 años)' })
    .max(120, { message: 'Edad no válida' }),
  
  email: z.string()
    .email({ message: 'Correo electrónico no válido' }),
  
  username: z.string()
    .min(4, { message: 'El usuario debe tener al menos 4 caracteres' })
    .max(20, { message: 'El usuario no puede tener más de 20 caracteres' })
    .regex(/^[a-zA-Z]+$/, { message: 'El usuario solo puede contener letras' }),
  
  password: z.string()
    .min(4, { message: 'La contraseña debe tener al menos 4 caracteres' })
    .max(20, { message: 'La contraseña no puede tener más de 20 caracteres' }),
  
  cpassword: z.string()
}).refine(data => data.password === data.cpassword, {
  message: 'Las contraseñas no coinciden',
  path: ['cpassword']
});