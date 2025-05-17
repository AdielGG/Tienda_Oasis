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


export const productSchema = z.object({
  nombre: z.string()
    .min(3, { message: 'El nombre debe tener al menos 3 caracteres' })
    .max(50, { message: 'El nombre no puede exceder 50 caracteres' }),
  
  descripcion: z.string()
    .min(10, { message: 'La descripción debe tener al menos 10 caracteres' })
    .max(500, { message: 'La descripción no puede exceder 500 caracteres' }),
    
  precio: z.number()
    .min(0.01, { message: 'El precio debe ser mayor a 0' })
    .max(10000, { message: 'El precio no puede exceder $10,000' }),
    
  cantidad: z.number()
    .int({ message: 'La cantidad debe ser un número entero' })
    .min(0, { message: 'La cantidad no puede ser negativa' })
    .max(1000, { message: 'La cantidad no puede exceder 1000 unidades' }),
    
  imagen: z.instanceof(File)
    .refine(file => file.size <= 2000000, { 
      message: 'La imagen no puede exceder 2MB' 
    })
    .refine(file => ['image/jpeg', 'image/png', 'image/bmp'].includes(file.type), {
      message: 'Solo se permiten imágenes JPEG, PNG o BMP'
    })
});

export const userSchema = z.object({
  nombre: z.string()
    .min(2, { message: 'El nombre debe tener al menos 2 caracteres' })
    .max(50, { message: 'El nombre no puede exceder 50 caracteres' }),
  
  lastname: z.string()
    .min(2, { message: 'El apellido debe tener al menos 2 caracteres' })
    .max(50, { message: 'El apellido no puede exceder 50 caracteres' }),
    
  usuario: z.string()
    .min(4, { message: 'El usuario debe tener al menos 4 caracteres' })
    .max(20, { message: 'El usuario no puede exceder 20 caracteres' })
    .regex(/^[a-zA-Z0-9_]+$/, { message: 'Solo letras, números y guiones bajos' }),
    
  correo: z.string()
    .email({ message: 'Correo electrónico no válido' }),
    
  rol: z.string()
    .refine(val => ['Administrador', 'Cliente', 'Dependiente', 'Soporte'].includes(val), {
      message: 'Rol no válido'
    }),
    
  contraseña: z.string()
    .min(6, { message: 'La contraseña debe tener al menos 6 caracteres' })
    .max(30, { message: 'La contraseña no puede exceder 30 caracteres' }),
    
  contraseñaConfirmacion: z.string()
}).refine(data => data.contraseña === data.contraseñaConfirmacion, {
  message: 'Las contraseñas no coinciden',
  path: ['contraseñaConfirmacion']
});