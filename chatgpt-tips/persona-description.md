##### The Key = Feed It and Seed IT – Feed copy; provide hints (Prompts & Parameters) to guide (Teach) the tool to succeed through your interactions and help you produce the effective content you want. 


https://5minutebi.com/2023/06/07/enhancing-technical-writing-with-chatgpt-a-step-by-step-prompt-engineering-approach/


### Prompt:
Acting as a [Role] Perform a [Task] using this [Guidance]

#### Prompt Structure & Roles / Personas

System Roles or a Persona in ChatGPT offer several benefits, allowing users to control the AI’s behavior and customize its responses. This also allows consistency not only between requests but between projects. Additionally, providing a Persona provides a more accessible and “relatable” interface.
Specify the Task to Perform

The “Task” refers to the specific action or objective the AI system is expected to perform. It defines the purpose or goal of the interaction. The task in prompt engineering can vary widely depending on the context and desired outcome. It could involve generating a creative story, answering specific questions, providing recommendations, offering explanations, or engaging in conversation on a particular topic.

By clearly defining the task in the prompt, technical writers can guide the AI system to produce responses that align with the intended purpose, leading to more effective and tailored user interactions.
Provide Guidance for the Task

#### Provide detailed context and background information through conversations with the AI.

This guidance serves as a direction for the AI to understand and perform the desired task effectively. The guidance helps shape the AI’s understanding and behavior, guiding it toward generating responses that align with the requirements to achieve your content goal.

Your target audience will have a specific skill level; ensure that your prompts include the specific technical language and terminology you want ChatGPT to use. This will help the model generate content regarding accuracy and technical depth.

By providing clear and concise guidance, prompt engineering ensures that the AI produces accurate and contextually appropriate outputs, enhancing the overall quality and relevance of its responses. Guidance can include examples such as;


    - Be explicit and detailed: The more specific your prompt, the better ChatGPT can generate the desired output.
    - Guide ChatGPT: Till the tool, what role are they to play? See example below
    - Set the format: If you need content in a particular structure, specify this in the prompt.
    - Define the tone and style: If you need the output to be specific, state this in the prompt.
    - Use instructive language: To guide the model toward generating content in a certain way, use language that provides clear instructions.
    - Technical Language to Include: Guide the prompts to include the specific technical language and terminology you want ChatGPT to use.
    - Iteration: Experiment with different prompts to see which yields the best results. Tell the chat what works, what does not, and why.
    - Context and Background: A detailed prompt allows the model to understand your requirements better and generate appropriate responses.


As an example, let’s take this use case. – “You are responsible for producing a technical how-to article into the hands of new users. You have a copy editor who will write the content, but you have to provide them with a content brief and outline. We will use ChatGPT Prompt Engineering Reference tables and the following Content Creation Framework to help with the first draft of the outline.”

<img decoding="async" src="https://i0.wp.com/5minutebi.com/wp-content/uploads/2023/06/ChatResponse14.png?resize=720%2C212&amp;ssl=1" alt="5MinuteBI Chat GPT Prompt Engineering Framework" class="wp-image-6602" srcset="https://i0.wp.com/5minutebi.com/wp-content/uploads/2023/06/ChatResponse14.png?w=720&amp;ssl=1 720w, https://i0.wp.com/5minutebi.com/wp-content/uploads/2023/06/ChatResponse14.png?resize=300%2C88&amp;ssl=1 300w" sizes="(max-width: 720px) 100vw, 720px" width="720" height="212">


### Step 0: Starting Point.

Let’s take this sample just as it is in a new ChatgGPT prompt. We will refine the content as we go. But first, let’s review how these results change as we move through the list.

Prompt: I need a content summary and outline about learning Power BI.


### Step 1 – Define the Persona of the AI (Identity)

In ChatGPT prompts, a persona is the language model’s assigned Role identity or character. It shapes the tone, style, and behavior of the generated responses. Creating a persona is like defining the AI assistant’s personality, allowing you to guide responses to align with a specific role or perspective. This can be a friendly customer support agent, expert, creative writer, or any desired persona.

To define a persona, you must specify their background, expertise, communication style, and personal traits. Create a persona as an experienced professional in the relevant industry for a formal and professional tone. For a casual and conversational tone, shape the Persona as a friendly and approachable companion. The Persona sets the context for AI responses and ensures consistency and coherence within the defined role between projects.

    Note: It’s important to remember that the persona is fictional. The AI model is a tool designed to generate text based on the persona and prompt.

This is a sample Persona I have been working on for a while. The words in the double brackets [] are ones I change depending on the project. I have a few Persona; Technical Writer, Python Programmer, and UX Designer.

Using this as a basis for a template helps put my lessons learned in one spot. See the ChatGPT Prompt Reference tables on the page for different parameters that can be used to define your inputs.

Prompt: ChatGPT adopts the following Persona or role with all of our interactions.

Persona Name = Jennifer Thompson

Background = [Jennifer] is an experienced [technical writer] with a strong background in [education]. They are passionate about creating [informative] and [engaging] content. With a deep understanding of [education theory and practices], [She] brings a unique perspective to their technical writing projects.

Expertise = [Jennifer] specializes in writing [technical documentation, blog posts, and lesson plans]. [She] has extensive knowledge of problem-based learning (PBL) approaches and incorporates this methodology into their instructional materials. They are skilled at breaking down complex concepts into clear and concise explanations, making technical information accessible to a wide range of audiences.  

Communication Style = [Jennifer] adopts a friendly and approachable communication style. They strive to create a conversational tone in their writing, ensuring that readers feel comfortable and engaged. With a focus on [clarity and simplicity], [Jennifer's] writing style helps readers grasp technical concepts while maintaining their interest throughout the content.

Personal Traits = [Jennifer] is highly detail-oriented, meticulous, and passionate about delivering [high-quality technical writing]. They possess [strong analytical skills and enjoy researching and staying updated on the latest industry trends]. [Jennifer] is also an effective problem solver and applies their problem-based learning approach to provide practical solutions and guidance through their writing.

[Skills Map]

Technical Writing Expertise = Identify Audience, Define Purpose, Research Topic, Draft Document, Revise & Edit 

Educational Background = Understanding Learning Theories, Developing Lesson Plans, Incorporating PBL, Tailoring Content for Different Learning Styles, Evaluating Student Progress 

Technology Skills = Power BI, Python, Excel, Data Visualization, Statistical Analysis, Predictive Modeling, Machine Learning, Big Data 

Communication Style = Friendly Tone, Engaging Writing, Focus on Clarity & Simplicity 

Personal Traits = Detail-Oriented, Meticulous, Passionate, Analytical, Problem Solver 



### Step 2 – Define Your Profile and Audience

If you have a specific brand you have developed, you can also design the Persona to align with your brand’s image, voice, and values. When you think about it, all your communications and products should help to reinforce your brand’s identity, and your AI assistant’s responses should help create a cohesive user experience.

You need to find your brand voice. Remember to think of the tool as more of an apprentice or assistant that can help you develop your content. As such, you need to define your business, your voice, and your target audience for this content.

Identifying the intended users of documentation, such as developers, end-users, or administrators, is essential. This allows for tailored content that addresses their needs and concerns, resulting in relevant and helpful documentation.

The following Mission Statement and Target Audience sections help to give clarity to you as an organization or workgroup. In addition, they allow you to help guide the AI to speak in your voice with some consistency between projects.

Prompt: I want our interactions to reflect following oranizational Mission Statement and Target Audience with our interactions.

Our Mission statement: Our educational program is dedicated to empowering individuals with the knowledge and skills required to harness the power of data analytics and data visualization using Power BI. Our mission is to transform people into data power users by providing concise, informative, and accessible content that can be consumed in short intervals. We emphasize our commitment to delivering digestible and efficient learning experiences for our audience.

Our Target Audience: Our overall content targets individuals new to data processing, analysis, visualization, data storytelling, and analytics. However, our target audience does have a basic knowledge of data analytics. We aim to provide valuable content for beginners and focus on how-to articles and tutorials centered around Power BI in addition to Machine Learning and analytics in Python.