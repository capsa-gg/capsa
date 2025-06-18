"use client";

import { yupResolver } from "@hookform/resolvers/yup";
import { Button, TextField } from "@mui/material";
import Box from "@mui/material/Box";
import type React from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";
import { useAddTitle, useGetAllTitles } from "@/api/hooks";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";
import { yupTitleEnvValidation } from "@/types/api/validation";

interface AddTitleForm {
    title: string;
}

const validationSchemaAddTitle = yup.object({
    title: yupTitleEnvValidation.title,
});

export const AddTitle: React.FC<AddTitleProps> = ({ onTitleAdded }) => {
    const { addNotification } = useNotificationsContext();
    const { trigger, isMutating } = useAddTitle();
    const { mutate } = useGetAllTitles();

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
        resetField,
    } = useForm<AddTitleForm>({
        resolver: yupResolver(validationSchemaAddTitle),
        mode: "onChange",
    });

    const onSubmit: SubmitHandler<AddTitleForm> = formData => {
        const title = formData.title;
        trigger(formData)
            .then(() => {
                resetField("title");
                addNotification({
                    severity: "success",
                    title: "Title added",
                    message: `The title ${formData.title} has been added successfully`,
                });
            })
            .finally(() => mutate().then(() => onTitleAdded(title)));
    };

    return (
        <Box
            component="form"
            noValidate
            onSubmit={handleSubmit(onSubmit)}
            sx={{
                display: "grid",
                columnGap: 2,
                gridTemplateColumns: "300px 100px",
                alignItems: "start",
                maxWidth: 800,
            }}
        >
            <TextField
                size="small"
                variant="outlined"
                required
                id="title"
                label="Title"
                error={Boolean(errors.title)}
                helperText={errors.title?.message}
                {...register("title")}
            />
            <Button disabled={!isValid || isMutating} type="submit" fullWidth variant="contained">
                Add
            </Button>
        </Box>
    );
};

export default AddTitle;

interface AddTitleProps {
    onTitleAdded: (value: string) => void;
}
