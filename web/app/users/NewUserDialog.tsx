import { yupResolver } from "@hookform/resolvers/yup";
import {
    Box,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    FormControl,
    InputLabel,
    MenuItem,
    Select,
    TextField,
} from "@mui/material";
import type React from "react";
import { Controller, type SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";
import useAdminUsers from "@/app/users/useAdminUsers";
import type { CreateUserRequest } from "@/types/api/users";
import { yupNewUserValidation } from "@/types/api/validation";

const validationSchema = yup.object().shape({
    email: yupNewUserValidation.email,
    firstName: yupNewUserValidation.firstName,
    lastName: yupNewUserValidation.lastName,
    role: yupNewUserValidation.role,
});

interface FormValues {
    email: string;
    firstName: string;
    lastName: string;
    role: string;
}

const NewUserDialog: React.FC<FormDialogProps> = ({ open, onClose }) => {
    const { allUsersIsLoading, isUpdating, addUser } = useAdminUsers();

    const {
        handleSubmit,
        control,
        reset,
        formState: { errors, isValid },
        register,
    } = useForm<FormValues>({
        resolver: yupResolver(validationSchema),
        mode: "onChange",
    });

    const onSubmit: SubmitHandler<FormValues> = formData => {
        addUser(formData as CreateUserRequest, () => reset()); // string enum doesn't work great with yup
        onClose();
    };

    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Add a new user</DialogTitle>
            <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
                <DialogContent>
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="email"
                        label="Email"
                        autoFocus
                        {...register("email")}
                        error={Boolean(errors.email)}
                        helperText={errors.email?.message}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="firstName"
                        label="First name"
                        {...register("firstName")}
                        error={Boolean(errors.firstName)}
                        helperText={errors.firstName?.message}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="lastName"
                        label="Last name"
                        {...register("lastName")}
                        error={Boolean(errors.lastName)}
                        helperText={errors.lastName?.message}
                    />
                    <FormControl fullWidth sx={{ mt: 3 }}>
                        <InputLabel>Role</InputLabel>
                        <Controller
                            name="role"
                            control={control}
                            defaultValue="User"
                            render={({ field }) => (
                                <Select {...field} label="Role">
                                    <MenuItem value="User">User</MenuItem>
                                    <MenuItem value="Admin">Admin</MenuItem>
                                </Select>
                            )}
                        />
                    </FormControl>
                </DialogContent>
                <DialogActions>
                    <Button onClick={onClose} color="primary">
                        Cancel
                    </Button>
                    <Button
                        variant="contained"
                        type="submit"
                        color="primary"
                        disabled={!isValid || allUsersIsLoading || isUpdating}
                    >
                        Create
                    </Button>
                </DialogActions>
            </Box>
        </Dialog>
    );
};

export default NewUserDialog;

interface FormDialogProps {
    open: boolean;
    onClose: () => void;
}
