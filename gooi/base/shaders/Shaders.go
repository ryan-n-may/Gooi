package shaders

const(
	VertexShaderSource = `
		#version 410

		layout(location = 0) in vec3 vp;
		layout(location = 1) in vec3 rgb;
		layout(location = 2) in vec3 sz;

		out vec3 colour;

		float l =  0.0f;
		float r =  sz.x;
		float b =  0.0f;
		float t =  sz.y;
		float n =  0.0f;
		float f =  1.0f;

		mat4 projection = mat4(
		    vec4(2.0/(r-l),     0.0,          0.0,         0.0),
		    vec4(0.0,           2.0/(t-b),    0.0,         0.0),
		    vec4(0.0,           0.0,         -2.0/(f-n),   0.0),
		    vec4(-(r+l)/(r-l), -(t+b)/(t-b), -(f+n)/(f-n), 1.0)
		);
		
		void main(void){
			gl_Position = projection*(vec4(vp.x, vp.y, vp.z, 1.0));
			colour = vec3(rgb.x, rgb.y, rgb.z);
		}
	` + "\x00"

	FragmentShaderSource = `
		#version 410

		in vec3 colour;
		out vec4 out_colour;

		void main(void) {
			out_colour = vec4(colour, 1.0);
		}
	` + "\x00"
)
