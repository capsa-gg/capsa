import * as yup from "yup";

export const yupAuthValidation = {
    email: yup.string().email("Enter a valid email").required("Email is required"),
    password: yup.string().min(16, "Password must be at least 16 characters").required("Password is required"),
    resetToken: yup.string().uuid("Reset token must be a valid UUID").required("Reset token is required"),
};

export const yupTitleEnvValidation = {
    title: yup
        .string()
        .min(3, "Minimum of 3 characters")
        .matches(/^[a-zA-Z0-9]+$/, "Only alphanumeric characters are allowed")
        .required("Title is required"),
    environment: yup
        .string()
        .min(3, "Minimum of 3 characters")
        .matches(/^[a-zA-Z0-9]+$/, "Only alphanumeric characters are allowed")
        .required("Environment is required"),
};

export const yupNewUserValidation = {
    email: yup.string().email("Enter a valid email").required("Email is required"),
    firstName: yup.string().min(2, "Minimum of 2 characters").required("First name is required"),
    lastName: yup.string().min(2, "Minimum of 2 characters").required("Last name is required"),
    role: yup.string().required("Role is required"),
};
